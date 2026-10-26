package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 bool; F1 int32; F2 uint16 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int8; F1 string }
type GenS6 struct { F0 uint16; F1 uint }
type GenS7 struct { F0 uint16; F1 uint; F2 int8 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 int8; F1 int8; F2 uint8 }
type GenS10 struct { F0 uint32; F1 int8 }
type GenS11 struct { F0 uint64; F1 uint64 }
type GenS12 struct { F0 uint32 }

func main() {
    var v1 float64 = 3.14159
    v2 := make(map[uint32]bool)
    v3 := make(map[string]bool)
    var v4 uint = 649
    v5 := []int{4294967295, -1, 1}
    for fi6 := 0; fi6 < 6; fi6++ {
        if (false || (!(32767 == 1))) {
            var v7 float32 = (1e10 + float32(1e10))
            _ = v7
        }
        _ = fi6
    }
    var v8 int = -678
    v9 := GenS1{F0: (false && (127 <= 128))}
    if ((v8 + 2147483647) > v8) {
        v4 = 4294967295
    } else {
        var v10 int = (v8 % ((-467) | 1))
        v8 -= 142
        _ = v10
    }
    if ("" != ("a" + "a")) {
        v1 += v1
        var v11 int32 = int32(v4)
        _ = len(v5)
        _ = v11
    } else {
        v4 -= v4
        switch v8 {
        case -406:
            v5 = append(v5, v8)
        case -1:
            for fi12 := 0; fi12 < 7; fi12++ {
                v4 = v4
                switch v8 {
                case -310:
                    var v13 int32 = 587
                    _ = v13
                case 32768:
                    var v14 string = "\n\t"
                    _ = v14
                default:
                    var v15 uint32 = 2
                    _ = v15
                }
                for fi16 := 0; fi16 < 13; fi16++ {
                    switch v8 {
                    case 42:
                        if true {
                            var v17 string = "a"
                            _ = v17
                        } else {
                            v18 := GenS2{F0: -2147483648}
                            fmt.Println(v4)
                            _ = v18
                        }
                    }
                    var v19 uint32 = 957
                    _ = fi16
                    _ = v19
                }
                _ = fi12
            }
        default:
            for fi20 := 0; fi20 < 17; fi20++ {
                if ((v4 | (v4 << (((v4 + (v4 + v4))) & 63))) <= (^(v4))) {
                    if (!false) {
                        var v21 bool = (65535 == 255)
                        if len(v5) > 0 {
                            v5[0] = -2147483648
                        }
                        _ = v21
                    } else {
                        var v22 int64 = 0
                        var v23 int8 = -128
                        switch (v8 * (-(fi20))) {
                        case 42:
                            if ((0 < 42) || ((((0.0 / (3.14159 - v1)) - float64((0.1 - 1.5))) == 1e10) && (!((("foo bar" >= "a") && (!(true || (!(true || ((v23 | (v23 * v23)) > (v23 >> (((v4 * (v4 | (v4 % ((v4) | 1))))) & 63)))))))) && false)))) {
                                if (!(32767 == -1)) {
                                    fmt.Println(v8)
                                    var v24 string = ("a" + "foo bar")
                                    _ = v24
                                }
                                fmt.Println(fi20)
                                for fi25 := 0; fi25 < 15; fi25++ {
                                    if true {
                                        _ = len(v3)
                                    } else {
                                        fmt.Println(v22, v8, v4)
                                        v1 = v1
                                    }
                                    _ = fi25
                                }
                            }
                        case 32767:
                            v26 := GenS3{F0: true, F1: 335, F2: 740}
                            _ = v26
                        default:
                            v27 := GenS4{F0: (^(v23))}
                            _ = v27
                        }
                        _ = v22
                        _ = v23
                    }
                    if false {
                        var v28 float64 = v1
                        v29 := GenS1{F0: (1 < 47)}
                        _ = v28
                        _ = v29
                    } else {
                        var v30 float32 = 0.1
                        for fi31 := 0; fi31 < 4; fi31++ {
                            if true {
                                fmt.Println(fi31, fi31, v30)
                                v1 = ((v1 - 1e10) * float64((0.1 / (float32((v1 * v1)) - (0.0 / (0.0 - 1.5))))))
                                fmt.Println(v8, v1, v8)
                            }
                            var v32 uint64 = 42
                            _ = fi31
                            _ = v32
                        }
                        _ = v30
                    }
                } else {
                    for fi33 := 0; fi33 < 5; fi33++ {
                        v34 := GenS5{F0: 100, F1: "a"}
                        _ = fi33
                        _ = v34
                    }
                }
                _ = fi20
            }
        }
        if (2 <= 4294967295) {
            fmt.Println(v8, v4, v1)
        }
    }
    switch "\n\t" {
    case "\n\t":
        var v35 uint32 = 760
        _ = v35
    case "hello":
        var v36 string = ""
        _ = v36
    default:
        v37 := GenS6{F0: 1, F1: (v4 % ((v4) | 1))}
        _ = v37
    }
    if (1 >= 9223372036854775807) {
        switch ("foo bar" + (("" + (("hello" + ("a" + (("a" + ((((("hello" + "\n\t") + "hello") + ("a" + "a")) + "foo bar") + "\n\t")) + ""))) + "a")) + "")) {
        case "":
            if (v4 >= v4) {
                fmt.Println(v1, v4)
            }
        case "a":
            v8 += v8
        default:
            if false {
                v1 = v1
            }
        }
    } else {
        if false {
            if false {
                if false {
                    v38 := GenS7{F0: 42, F1: 4294967295, F2: -2}
                    if (!(1 > 0)) {
                        v1 += v1
                        for fi39 := 0; fi39 < 5; fi39++ {
                            var v40 bool = (514 == -32768)
                            if (32767 > -396) {
                                if (((1.5 * 1e10) / float32(v1)) >= (((3.14159 + ((float32(v1) + 0.1) + 0.0)) - 0.0) * (((1e10 - 1.5) * 1e10) + ((0.1 / (1e10 / 1e10)) + float32(v1))))) {
                                    for fi41 := 0; fi41 < 8; fi41++ {
                                        v5 = append(v5, fi41)
                                        delete(v3, "foo bar")
                                        fmt.Println(fi41)
                                        _ = fi41
                                    }
                                    v40 = (v4 > v4)
                                } else {
                                    v42 := GenS8{F0: 1}
                                    _ = v42
                                }
                                v43 := GenS4{F0: -7}
                                _ = v43
                            } else {
                                v40 = false
                                fmt.Println(v40, fi39, v1)
                            }
                            v8 = 32767
                            _ = fi39
                            _ = v40
                        }
                        var v44 int64 = -9223372036854775808
                        _ = v44
                    }
                    _ = v38
                } else {
                    if (!(("foo bar" + "a") != "foo bar")) {
                        switch "\n\t" {
                        case "hello":
                            if true {
                                for fi45 := 0; fi45 < 15; fi45++ {
                                    var v46 uint8 = 136
                                    for fi47 := 0; fi47 < 12; fi47++ {
                                        v8 = (fi45 ^ (v8 & (^(v8))))
                                        v48 := GenS4{F0: -1}
                                        var v49 uint16 = 669
                                        _ = fi47
                                        _ = v48
                                        _ = v49
                                    }
                                    fmt.Println(v46)
                                    _ = fi45
                                    _ = v46
                                }
                                if true {
                                    var v50 int64 = int64(v8)
                                    switch "" {
                                    case "a":
                                        var v51 int8 = 127
                                        _ = v51
                                    }
                                    fmt.Println(v4, v8, v50)
                                    _ = v50
                                } else {
                                    v1 = ((0.1 * (v1 / v1)) * -2.25)
                                }
                                fmt.Println(v8)
                            } else {
                                switch "foo bar" {
                                case "a":
                                    v1 += v1
                                case "\n\t":
                                    v4 += (v4 / ((18446744073709551615) | 1))
                                default:
                                    v4 = 1
                                }
                            }
                        case "":
                            delete(v3, ("\n\t" + "\n\t"))
                        default:
                            v8 -= (v8 & v8)
                        }
                        if true {
                            fmt.Println(v4, v1, v1)
                        }
                    } else {
                        fmt.Println(v4, v1, v1)
                    }
                }
            } else {
                _ = len(v5)
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                v8 = (v8 * -1)
            }
            v52 := GenS9{F0: 100, F1: 99, F2: 111}
            v1 = v1
            _ = v52
        }
        var v53 int = 2147483647
        fmt.Println(v53, v53)
        _ = v53
    }
    v54 := GenS10{F0: 533, F1: -1}
    v1 = 3.14159
    fmt.Println(v1, v8, v4)
    var v55 uint = v4
    for fi56 := 0; fi56 < 10; fi56++ {
        var v57 int = (^(v8))
        v58 := v2[1]
        _ = fi56
        _ = v57
        _ = v58
    }
    v2[815] = (834 >= 371)
    if true {
        fmt.Println(v8, v8)
    } else {
        for fi59 := 0; fi59 < 16; fi59++ {
            for fi60 := 0; fi60 < 1; fi60++ {
                v61 := GenS6{F0: 614, F1: 0}
                _ = fi60
                _ = v61
            }
            v8 = v8
            if (32767 != 552) {
                for fi62 := 0; fi62 < 14; fi62++ {
                    var v63 int64 = -1
                    _ = fi62
                    _ = v63
                }
            }
            _ = fi59
        }
    }
    if false {
        switch 1 {
        case 32768:
            var v64 string = ("" + "hello")
            _ = v64
        case -380:
            v65 := GenS11{F0: 0, F1: 2}
            _ = v65
        }
    } else {
        v66 := GenS4{F0: -51}
        if false {
            for fi67 := 0; fi67 < 19; fi67++ {
                v1 -= (v1 + v1)
                _ = fi67
            }
            _ = len(v2)
        } else {
            if (255 == 201) {
                var v68 int16 = -441
                delete(v3, ("a" + ("a" + "foo bar")))
                v69 := GenS9{F0: int8(v4), F1: 110, F2: 59}
                _ = v68
                _ = v69
            } else {
                v8 = (v8 * v8)
                var v70 float64 = 3.14159
                _ = v70
            }
            v1 += (1e10 * float64(1e10))
            if false {
                v8 += v8
            }
        }
        v71 := v2[65535]
        _ = v66
        _ = v71
    }
    v8 -= v8
    switch "hello" {
    case "\n\t":
        for fi72 := 0; fi72 < 6; fi72++ {
            fmt.Println(v8, v1, v4)
            v8 += 2147483647
            v73 := GenS12{F0: 2147483647}
            _ = fi72
            _ = v73
        }
    case "hello":
        var v74 uint64 = 4294967295
        _ = v74
    case "foo bar":
        v75 := v2[2147483647]
        _ = v75
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v55:", v55)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v55
    _ = v8
    _ = v9
}
