package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32; F1 int16; F2 bool }
type GenS3 struct { F0 int32; F1 string }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint; F1 int; F2 int8 }
type GenS6 struct { F0 uint16; F1 int64; F2 int64 }
type GenS7 struct { F0 string; F1 int64; F2 int16 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 uint64; F1 uint32; F2 uint32 }
type GenS10 struct { F0 int32; F1 int64 }
type GenS11 struct { F0 int16; F1 int }

func main() {
    var v1 bool = false
    var v2 uint64 = 4294967295
    v3 := []int{-2147483648, 32768, 9223372036854775807, 2147483647}
    var v4 uint = 450
    if (true && true) {
        fmt.Println(v1, v1, v1)
        switch "a" {
        case "\n\t":
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        case "":
            _ = len(v3)
        case "hello":
            v5 := GenS1{F0: v2}
            _ = v5
        }
        if (float32(1e10) <= (0.0 - 3.14159)) {
            v6 := GenS2{F0: 2147483648, F1: -147, F2: true}
            var v7 uint8 = 128
            if v1 {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
            _ = v6
            _ = v7
        }
    } else {
        var v8 int32 = -531
        _ = v8
    }
    v9 := GenS3{F0: 42, F1: (((("" + "") + "\n\t") + ("a" + "a")) + "a")}
    fmt.Println(v1)
    fmt.Println(v2, v1, v2)
    v10 := GenS4{F0: -1}
    var v11 int = -198
    var v12 int = v11
    switch ("" + "foo bar") {
    case "\n\t":
        if false {
            if v1 {
                var v13 uint = v4
                for fi14 := 0; fi14 < 13; fi14++ {
                    var v15 uint16 = 235
                    var v16 uint16 = (v15 * 42)
                    _ = fi14
                    _ = v15
                    _ = v16
                }
                _ = v13
            } else {
                if true {
                    if (!((("a" + (("" + "\n\t") + "\n\t")) + "a") != "hello")) {
                        v2 = v2
                    }
                    v4 = (^(v4))
                } else {
                    v17 := GenS5{F0: v4, F1: (v12 << ((v4) & 63)), F2: -34}
                    _ = v17
                }
                switch v11 {
                case 127:
                    var v18 int32 = 42
                    _ = v18
                case 932:
                    var v19 uint8 = 11
                    _ = v19
                }
            }
        } else {
            for fi20 := 0; fi20 < 6; fi20++ {
                var v21 uint32 = 2147483648
                fmt.Println(v2, v11)
                v2 -= v2
                _ = fi20
                _ = v21
            }
            v2 += (v2 << ((v4) & 63))
            if v1 {
                if v1 {
                    v2 = v2
                } else {
                    v4 = v4
                }
                v22 := GenS4{F0: 2147483647}
                var v23 int16 = -195
                _ = v22
                _ = v23
            } else {
                for fi24 := 0; fi24 < 3; fi24++ {
                    var v25 uint8 = 255
                    v25 = (^(v25))
                    _ = fi24
                    _ = v25
                }
                for fi26 := 0; fi26 < 15; fi26++ {
                    var v27 int32 = 1
                    v27 += (^(v27))
                    _ = fi26
                    _ = v27
                }
                for fi28 := 0; fi28 < 16; fi28++ {
                    var v29 float32 = 0.0
                    _ = fi28
                    _ = v29
                }
            }
        }
    case "foo bar":
        for fi30 := 0; fi30 < 12; fi30++ {
            if (!((v4 > 255) || (-116 > 42))) {
                for fi31 := 0; fi31 < 11; fi31++ {
                    var v32 int8 = -1
                    if (661 <= -222) {
                        v32 = (-(v32))
                        var v33 uint8 = 1
                        v12 -= v12
                        _ = v33
                    }
                    _ = fi31
                    _ = v32
                }
                v1 = (65535 > 312)
                v34 := GenS6{F0: uint16(v2), F1: 78, F2: 42}
                _ = v34
            }
            if (2147483647 == 42) {
                switch "\n\t" {
                case "foo bar":
                    v11 -= (^(fi30))
                case "a":
                    if (2147483648 != 2147483647) {
                        if v1 {
                            var v35 int32 = int32(fi30)
                            var v36 string = "foo bar"
                            _ = v35
                            _ = v36
                        } else {
                            var v37 uint8 = 0
                            var v38 uint8 = v37
                            var v39 uint16 = 1
                            _ = v37
                            _ = v38
                            _ = v39
                        }
                        _ = len(v3)
                    } else {
                        switch (v12 << ((42) & 63)) {
                        case -246:
                            for fi40 := 0; fi40 < 16; fi40++ {
                                v41 := GenS4{F0: -176}
                                if v1 {
                                    v2 += uint64(v4)
                                    var v42 bool = (!true)
                                    if (true || false) {
                                        if v1 {
                                            var v43 string = "\n\t"
                                            v1 = (392 < 0)
                                            var v44 string = ("" + (v43 + ("a" + v43)))
                                            _ = v43
                                            _ = v44
                                        }
                                        var v45 float32 = (0.0 * 1e10)
                                        _ = v45
                                    } else {
                                        var v46 float64 = (0.0 / (0.1 + (1.5 * 0.1)))
                                        v3 = append(v3, (^(fi30)))
                                        _ = v46
                                    }
                                    _ = v42
                                }
                                _ = fi40
                                _ = v41
                            }
                        case -9223372036854775808:
                            var v47 string = ("\n\t" + "a")
                            _ = v47
                        default:
                            v48 := GenS7{F0: "foo bar", F1: -623, F2: 1000}
                            _ = v48
                        }
                        fmt.Println(fi30, v4)
                    }
                default:
                    var v49 uint8 = 127
                    _ = v49
                }
                var v50 uint = v4
                _ = v50
            } else {
                v11 = 42
                v4 = v4
            }
            var v51 string = ("foo bar" + "a")
            _ = fi30
            _ = v51
        }
    }
    if (0 > 255) {
        var v52 int = 509
        var v53 int16 = 42
        _ = v52
        _ = v53
    }
    if false {
        fmt.Println(v11, v1)
        for fi54 := 0; fi54 < 18; fi54++ {
            v55 := GenS8{F0: 2}
            _ = fi54
            _ = v55
        }
        var v56 float64 = -2.25
        _ = v56
    } else {
        v3 = append(v3, v11)
        for fi57 := 0; fi57 < 2; fi57++ {
            switch "" {
            case "foo bar":
                switch 4294967295 {
                case 99:
                    v4 = v4
                case 42:
                    if v1 {
                        v11 += (fi57 + (v11 ^ (-(fi57))))
                    }
                case 4294967295:
                    v58 := GenS3{F0: -2147483648, F1: "hello"}
                    _ = v58
                }
            case "":
                var v59 float64 = 0.1
                _ = v59
            case "hello":
                fmt.Println(v4)
            }
            _ = fi57
        }
    }
    fmt.Println(v11, v11, v4)
    if v1 {
        v60 := GenS9{F0: 872, F1: 65535, F2: 42}
        _ = v60
    } else {
        switch v12 {
        case -502:
            if len(v3) > 0 {
                v3[0] = v12
            }
        case 16:
            var v61 uint16 = 1000
            _ = v61
        case -2147483648:
            switch "foo bar" {
            case "a":
                v11 -= 350
            case "foo bar":
                if len(v3) > 0 {
                    v3[0] = (v12 * v12)
                }
            }
        default:
            var v62 float64 = (float64(((1.5 / (1e10 + 0.1)) / -2.25)) / (float64(((0.0 - 3.14159) / (((3.14159 * float32((-2.25 / 1e10))) / 1.0) + 3.14159))) - ((0.1 * 1.5) - (3.14159 + 3.14159))))
            _ = v62
        }
    }
    v12 -= v12
    v2 -= v2
    for fi63 := 0; fi63 < 17; fi63++ {
        if true {
            var v64 uint = uint(v2)
            var v65 int64 = -9223372036854775808
            if v1 {
                fmt.Println(v4, v65)
            } else {
                for fi66 := 0; fi66 < 8; fi66++ {
                    v11 -= (v11 ^ v12)
                    v65 += v65
                    _ = fi66
                }
                if (v1 && v1) {
                    if v1 {
                        var v67 uint16 = 1
                        _ = v67
                    } else {
                        for fi68 := 0; fi68 < 8; fi68++ {
                            switch (-(fi68)) {
                            case 393:
                                v69 := GenS10{F0: -1, F1: v65}
                                _ = v69
                            }
                            switch fi68 {
                            case -9223372036854775808:
                                var v70 uint32 = 4294967295
                                _ = v70
                            }
                            _ = fi68
                        }
                        fmt.Println(v65)
                    }
                    switch -9223372036854775808 {
                    case -9223372036854775808:
                        for fi71 := 0; fi71 < 2; fi71++ {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            v1 = (v1 && v1)
                            _ = fi71
                        }
                    default:
                        _ = len(v3)
                    }
                    v72 := GenS9{F0: v2, F1: 65535, F2: 42}
                    _ = v72
                } else {
                    fmt.Println(v1)
                    v65 = -1
                    v73 := GenS5{F0: (v4 - v64), F1: v12, F2: 127}
                    _ = v73
                }
            }
            _ = v64
            _ = v65
        } else {
            v12 += (v12 / ((1) | 1))
            switch "foo bar" {
            case "":
                v4 += (v4 << (((^(v4))) & 63))
            case "\n\t":
                switch "\n\t" {
                case "foo bar":
                    if false {
                        v74 := GenS11{F0: 32767, F1: v11}
                        _ = v74
                    } else {
                        v1 = ((!(1e10 < 1.5)) && false)
                    }
                }
            }
            v3 = append(v3, v11)
        }
        _ = fi63
    }
    v75 := GenS9{F0: v2, F1: 0, F2: uint32(v2)}
    fmt.Println(v11, v12, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v75
    _ = v9
}
