package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 uint }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int16; F1 uint16 }
type GenS4 struct { F0 int; F1 uint16; F2 string }
type GenS5 struct { F0 int16; F1 bool; F2 int32 }
type GenS6 struct { F0 bool; F1 int; F2 uint64 }
type GenS7 struct { F0 uint }

func main() {
    var v1 int64 = 16
    v2 := []string{""}
    var v3 float64 = 0.1
    var v4 float32 = 0.0
    var v5 int64 = 1
    fmt.Println(v1, v5, v5)
    var v6 uint8 = 255
    v1 = (v1 + v1)
    var v7 int16 = int16(v1)
    var v8 uint16 = 1
    if true {
        for fi9 := 0; fi9 < 7; fi9++ {
            for fi10 := 0; fi10 < 15; fi10++ {
                v11 := GenS1{F0: (("" + ("foo bar" + "a")) + "a"), F1: int8(v1), F2: 1}
                var v12 float64 = v3
                v13 := GenS2{F0: 47}
                _ = fi10
                _ = v11
                _ = v12
                _ = v13
            }
            for fi14 := 0; fi14 < 17; fi14++ {
                fmt.Println(v6, v7)
                _ = fi14
            }
            var v15 bool = true
            _ = fi9
            _ = v15
        }
    }
    for fi16 := 0; fi16 < 6; fi16++ {
        switch "foo bar" {
        case "hello":
            if len(v2) > 0 {
                v2[0] = "\n\t"
            }
        case "foo bar":
            for fi17 := 0; fi17 < 13; fi17++ {
                v8 += v8
                _ = fi17
            }
        case "":
            v1 = 9223372036854775807
        }
        switch fi16 {
        case 2147483647:
            var v18 uint64 = 42
            _ = v18
        default:
            for fi19 := 0; fi19 < 15; fi19++ {
                v20 := GenS2{F0: int8(fi16)}
                _ = fi19
                _ = v20
            }
        }
        var v21 int8 = 76
        _ = fi16
        _ = v21
    }
    v3 = 1e10
    v1 = (v1 * v5)
    fmt.Println(v8)
    v2 = append(v2, "")
    var v22 int16 = v7
    switch ("foo bar" + "\n\t") {
    case "\n\t":
        fmt.Println(v8, v3)
    case "foo bar":
        switch ((((("hello" + "foo bar") + ("" + "a")) + "foo bar") + "foo bar") + "hello") {
        case "hello":
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
    default:
        switch "hello" {
        case "":
            switch ("foo bar" + (("foo bar" + ("hello" + "\n\t")) + "foo bar")) {
            case "hello":
                v3 = 3.14159
            case "foo bar":
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            case "a":
                var v23 float64 = v3
                _ = v23
            default:
                var v24 float32 = 0.0
                _ = v24
            }
        case "\n\t":
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
    }
    var v25 int64 = v1
    if false {
        if false {
            v1 -= v1
            for fi26 := 0; fi26 < 9; fi26++ {
                var v27 int64 = (-(v1))
                v28 := GenS1{F0: "\n\t", F1: 1, F2: 0}
                _ = fi26
                _ = v27
                _ = v28
            }
        }
        switch "a" {
        case "\n\t":
            if true {
                for fi29 := 0; fi29 < 1; fi29++ {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    if (3.14159 >= float64(v4)) {
                        for fi30 := 0; fi30 < 10; fi30++ {
                            if (-1 < -128) {
                                switch ("\n\t" + ("a" + "a")) {
                                case "hello":
                                    if false {
                                        switch (fi30 ^ fi29) {
                                        case -727:
                                            v31 := GenS1{F0: "", F1: 87, F2: 4294967295}
                                            _ = v31
                                        case 2147483647:
                                            var v32 string = "hello"
                                            _ = v32
                                        default:
                                            var v33 int32 = 1
                                            v25 += (v5 ^ 0)
                                            _ = v33
                                        }
                                        fmt.Println(v6, v22, fi29)
                                        var v34 string = "foo bar"
                                        _ = v34
                                    } else {
                                        v6 -= (v6 ^ (v6 >> ((255) & 63)))
                                        v8 -= (v8 / ((42) | 1))
                                    }
                                }
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                if true {
                                    v4 = (v4 / v4)
                                    v5 += v25
                                } else {
                                    var v35 int32 = -2147483648
                                    if (!((fi30 / (((fi29 - (^(fi29)))) | 1)) < (-(fi29)))) {
                                        var v36 float64 = ((v3 * float64((((3.14159 + (-2.25 + v4)) / 3.14159) + v4))) - v3)
                                        v8 -= 32767
                                        _ = v36
                                    } else {
                                        v6 = v6
                                    }
                                    v8 = v8
                                    _ = v35
                                }
                                var v37 int64 = 367
                                switch 1 {
                                case -9223372036854775808:
                                    if false {
                                        v38 := GenS3{F0: 1, F1: (v8 % ((v8) | 1))}
                                        var v39 int64 = v5
                                        if false {
                                            if false {
                                                var v40 int32 = 1
                                                var v41 string = "hello"
                                                _ = v40
                                                _ = v41
                                            }
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        } else {
                                            v37 = v5
                                            fmt.Println(v39, v5)
                                        }
                                        _ = v38
                                        _ = v39
                                    }
                                    if len(v2) > 0 {
                                        v2[0] = ""
                                    }
                                    fmt.Println(v7, fi30)
                                }
                                v37 = 42
                                _ = v37
                            } else {
                                var v42 uint64 = 42
                                _ = v42
                            }
                            _ = len(v2)
                            var v43 float32 = (0.0 + 1.5)
                            v43 -= v43
                            fmt.Println(v8)
                            _ = fi30
                            _ = v43
                        }
                        if false {
                            var v44 string = ("hello" + ("\n\t" + "hello"))
                            for fi45 := 0; fi45 < 5; fi45++ {
                                if true {
                                    if (414 <= 197) {
                                        v46 := GenS4{F0: 32767, F1: uint16(v6), F2: "hello"}
                                        v2 = append(v2, v44)
                                        if true {
                                            switch fi29 {
                                            case 0:
                                                var v47 uint32 = 42
                                                _ = v47
                                            case -846:
                                                v25 = (v25 ^ (v5 >> ((0) & 63)))
                                            default:
                                                var v48 uint8 = (v6 & (v6 << ((1) & 63)))
                                                _ = v48
                                            }
                                            var v49 bool = true
                                            _ = v49
                                        } else {
                                            if true {
                                                fmt.Println(v22, v7, v3)
                                                fmt.Println(v25, fi29)
                                                v4 += (v4 - v4)
                                            }
                                            var v50 float64 = float64(float32((v3 * v3)))
                                            fmt.Println(v50, v1, v50)
                                            _ = v50
                                        }
                                        _ = v46
                                    }
                                } else {
                                    var v51 uint64 = 4294967295
                                    _ = v51
                                }
                                var v52 int16 = (v7 / ((1000) | 1))
                                _ = fi45
                                _ = v52
                            }
                            _ = v44
                        } else {
                            var v53 int8 = int8(v6)
                            _ = v53
                        }
                        var v54 string = ("\n\t" + "hello")
                        _ = v54
                    } else {
                        for fi55 := 0; fi55 < 14; fi55++ {
                            v6 += (^(v6))
                            v7 += (v22 ^ v22)
                            fmt.Println(fi55)
                            _ = fi55
                        }
                        fmt.Println(v1, v22, v3)
                    }
                    _ = fi29
                }
                _ = len(v2)
            } else {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v56 := GenS3{F0: -32768, F1: 2}
                _ = v56
            }
        case "foo bar":
            if ((1 <= 42) || (0 > 707)) {
                switch 768 {
                case 32767:
                    for fi57 := 0; fi57 < 3; fi57++ {
                        var v58 uint64 = 4294967295
                        switch int(v6) {
                        case 156:
                            if false {
                                switch fi57 {
                                case 1:
                                    if true {
                                        v59 := GenS5{F0: (v22 >> ((255) & 63)), F1: (42 <= 13), F2: -1}
                                        if (true || (!true)) {
                                            v7 -= (v22 + (v7 << ((1) & 63)))
                                            _ = len(v2)
                                        }
                                        _ = v59
                                    } else {
                                        switch "a" {
                                        case "\n\t":
                                            fmt.Println(fi57)
                                        case "":
                                            fmt.Println(fi57)
                                        default:
                                            v5 = v25
                                        }
                                        v60 := GenS6{F0: false, F1: -669, F2: (v58 >> ((255) & 63))}
                                        _ = v60
                                    }
                                case 32767:
                                    fmt.Println(v5, v1)
                                }
                                v25 = 42
                            } else {
                                for fi61 := 0; fi61 < 16; fi61++ {
                                    var v62 float32 = (v4 + 0.1)
                                    _ = fi61
                                    _ = v62
                                }
                                if true {
                                    if ((!false) && true) {
                                        v63 := GenS2{F0: 97}
                                        v22 -= int16(v6)
                                        _ = v63
                                    } else {
                                        var v64 uint32 = 2
                                        for fi65 := 0; fi65 < 2; fi65++ {
                                            v66 := GenS3{F0: v22, F1: v8}
                                            _ = fi65
                                            _ = v66
                                        }
                                        _ = v64
                                    }
                                } else {
                                    fmt.Println(v5, v7)
                                    v22 += (v22 / (((^(v7))) | 1))
                                }
                                v67 := GenS7{F0: 0}
                                _ = v67
                            }
                        default:
                            var v68 bool = (1 < 4294967295)
                            _ = v68
                        }
                        v4 += ((v4 * -2.25) - 3.14159)
                        _ = fi57
                        _ = v58
                    }
                default:
                    v2 = append(v2, "")
                }
            }
        case "a":
            v69 := GenS6{F0: (!(v1 != (v5 | v25))), F1: 366, F2: 42}
            _ = v69
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v22
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
