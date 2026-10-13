package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 int64; F1 uint64; F2 uint32 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint; F1 uint64; F2 bool }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int8; F1 uint64; F2 uint64 }
type GenS7 struct { F0 int16; F1 int16 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 bool; F1 int64 }
type GenS10 struct { F0 uint32; F1 string; F2 uint32 }
type GenS11 struct { F0 int64; F1 uint8 }
type GenS12 struct { F0 uint16; F1 uint }
type GenS13 struct { F0 int8 }

func main() {
    v1 := make(map[string]string)
    var v2 uint16 = 42
    var v3 int64 = 1
    v4 := make(map[string]bool)
    v5 := []uint32{}
    var v6 int32 = -244
    v7 := GenS1{F0: 42, F1: (v6 << ((78) & 63))}
    var v8 int64 = 0
    if true {
        var v9 uint64 = 42
        _ = v9
    }
    switch "hello" {
    case "":
        for fi10 := 0; fi10 < 4; fi10++ {
            fmt.Println(v2)
            var v11 uint8 = 61
            if (true || ((("\n\t" != ("\n\t" + "")) || (-34 == 127)) || (true && false))) {
                v12 := GenS1{F0: uint(v2), F1: -268}
                v13 := GenS2{F0: v3, F1: 847, F2: 2147483648}
                if (((^(v3)) >= (v8 * -1)) || true) {
                    if (v11 < (v11 & 1)) {
                        _ = len(v1)
                        _ = len(v1)
                        var v14 int32 = -1
                        _ = v14
                    } else {
                        v15 := GenS3{F0: 0}
                        fmt.Println(v11, v6)
                        _ = len(v5)
                        _ = v15
                    }
                    switch -1 {
                    case 0:
                        switch 698 {
                        case 2147483647:
                            switch "hello" {
                            case "\n\t":
                                if false {
                                    switch ((("foo bar" + "") + ("foo bar" + ("hello" + "foo bar"))) + "hello") {
                                    case "\n\t":
                                        if true {
                                            fmt.Println(v8, v3)
                                            for fi16 := 0; fi16 < 11; fi16++ {
                                                v4["hello"] = false
                                                var v17 string = ""
                                                var v18 float32 = 0.0
                                                _ = fi16
                                                _ = v17
                                                _ = v18
                                            }
                                            _ = len(v5)
                                        } else {
                                            v6 -= int32(fi10)
                                            if false {
                                                fmt.Println(v8, v6)
                                            } else {
                                                v19 := GenS4{F0: 410, F1: 18446744073709551615, F2: (!(true || (!(18446744073709551615 > 42))))}
                                                var v20 int = (fi10 & (fi10 + (^(fi10))))
                                                _ = v19
                                                _ = v20
                                            }
                                        }
                                    case "":
                                        var v21 uint = 18446744073709551615
                                        _ = v21
                                    case "hello":
                                        if (v6 >= -1) {
                                            var v22 bool = (!(false && (v2 < (v2 % ((v2) | 1)))))
                                            var v23 uint32 = 942
                                            _ = v22
                                            _ = v23
                                        } else {
                                            var v24 uint8 = 255
                                            v25 := GenS1{F0: 255, F1: 315}
                                            v6 += (v6 - v6)
                                            _ = v24
                                            _ = v25
                                        }
                                    default:
                                        v3 -= v3
                                    }
                                    if (false && (!false)) {
                                        v26 := GenS2{F0: v3, F1: 249, F2: 841}
                                        _ = v26
                                    } else {
                                        for fi27 := 0; fi27 < 16; fi27++ {
                                            v2 += 42
                                            fmt.Println(fi10, v8)
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                            _ = fi27
                                        }
                                    }
                                    v28 := GenS2{F0: v8, F1: 979, F2: uint32(v11)}
                                    _ = v28
                                }
                            case "a":
                                v29 := GenS5{F0: -502}
                                _ = v29
                            default:
                                for fi30 := 0; fi30 < 1; fi30++ {
                                    if (!(-2.25 <= (3.14159 + ((3.14159 / 0.1) - float32(0.0))))) {
                                        switch (fi10 & 0) {
                                        case 1:
                                            fmt.Println(v2)
                                        case -9223372036854775808:
                                            var v31 uint16 = v2
                                            _ = v31
                                        }
                                        var v32 int8 = 63
                                        if true {
                                            _ = len(v4)
                                            fmt.Println(v6, v6, v6)
                                            _ = len(v5)
                                        } else {
                                            if false {
                                                fmt.Println(v3, v6, v3)
                                                v6 = (v6 + (^(v6)))
                                            }
                                        }
                                        _ = v32
                                    } else {
                                        v33 := GenS5{F0: -32768}
                                        v34 := GenS6{F0: 119, F1: 0, F2: 140}
                                        v6 = (v6 * v6)
                                        _ = v33
                                        _ = v34
                                    }
                                    v35 := GenS7{F0: 1, F1: 32767}
                                    _ = fi30
                                    _ = v35
                                }
                            }
                        case 829:
                            switch "hello" {
                            case "foo bar":
                                fmt.Println(v11, v6)
                            case "hello":
                                if len(v5) > 0 {
                                    v5[0] = 2147483647
                                }
                            case "":
                                v3 = 9223372036854775807
                            default:
                                var v36 int = -328
                                _ = v36
                            }
                        default:
                            if ((!(true || (((float64(-2.25) * (-2.25 + 0.0)) <= (((0.0 - float64(0.1)) + 3.14159) - 3.14159)) && (663 > 1000)))) && true) {
                                if ((!(v6 < v6)) && true) {
                                    var v37 string = ("" + "hello")
                                    v6 -= int32(v11)
                                    var v38 uint8 = v11
                                    _ = v37
                                    _ = v38
                                } else {
                                    v11 -= 2
                                    v39 := GenS8{F0: -761}
                                    v3 += (v3 * 574)
                                    _ = v39
                                }
                            }
                        }
                    case 1:
                        if (!false) {
                            for fi40 := 0; fi40 < 12; fi40++ {
                                if (int(v3) == fi40) {
                                    if false {
                                        if (142 <= 42) {
                                            v8 = (v8 * (v8 + (^(v8))))
                                        } else {
                                            switch "\n\t" {
                                            case "":
                                                v8 += -304
                                                switch "" {
                                                case "":
                                                    if (669 >= 2147483648) {
                                                        v41 := v1[(("a" + "\n\t") + ("foo bar" + (("" + "foo bar") + ("" + "hello"))))]
                                                        _ = v41
                                                    } else {
                                                        fmt.Println(v3, v2)
                                                        var v42 int8 = 1
                                                        _ = v42
                                                    }
                                                case "foo bar":
                                                    switch (fi40 & (fi40 * (fi40 % (((fi10 / (((fi10 & (fi40 & fi40))) | 1))) | 1)))) {
                                                    case -146:
                                                        v6 = (v6 << ((1) & 63))
                                                        var v43 string = ""
                                                        var v44 bool = true
                                                        _ = v43
                                                        _ = v44
                                                    }
                                                case "a":
                                                    v2 -= 255
                                                default:
                                                    fmt.Println(v6)
                                                }
                                                v8 += v3
                                            case "a":
                                                v45 := GenS6{F0: 100, F1: 902, F2: 323}
                                                _ = v45
                                            case "foo bar":
                                                v11 = v11
                                            default:
                                                v46 := GenS9{F0: false, F1: (-(v8))}
                                                _ = v46
                                            }
                                        }
                                    } else {
                                        v8 -= v8
                                        var v47 string = "a"
                                        _ = v47
                                    }
                                    if ((true && (-2.25 == -2.25)) && true) {
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                    } else {
                                        var v48 uint = 636
                                        fmt.Println(v48, v48)
                                        _ = v48
                                    }
                                }
                                delete(v1, (((((("\n\t" + "") + "") + ("foo bar" + ("" + "foo bar"))) + (("\n\t" + "") + (("\n\t" + "hello") + "a"))) + "hello") + "foo bar"))
                                _ = fi40
                            }
                        } else {
                            var v49 bool = false
                            v50 := GenS10{F0: 2147483648, F1: "\n\t", F2: 0}
                            _ = v49
                            _ = v50
                        }
                    }
                    v51 := GenS9{F0: ((false && true) && true), F1: (v8 << ((65535) & 63))}
                    _ = v51
                }
                _ = v12
                _ = v13
            } else {
                var v52 int64 = v8
                _ = v52
            }
            _ = fi10
            _ = v11
        }
    case "\n\t":
        fmt.Println(v6, v6)
    case "a":
        v3 -= v8
    }
    switch -952 {
    case -572:
        for fi53 := 0; fi53 < 13; fi53++ {
            for fi54 := 0; fi54 < 14; fi54++ {
                v55 := GenS11{F0: 42, F1: 2}
                fmt.Println(v2)
                var v56 uint32 = 2147483648
                _ = fi54
                _ = v55
                _ = v56
            }
            var v57 float64 = (0.0 * 3.14159)
            _ = fi53
            _ = v57
        }
    case -2:
        v58 := GenS5{F0: -1}
        _ = v58
    }
    var v59 uint64 = 34
    var v60 int8 = -128
    v61 := v1["a"]
    for fi62 := 0; fi62 < 19; fi62++ {
        if (int64(v60) > v3) {
            if false {
                v63 := GenS3{F0: 2}
                fmt.Println(fi62)
                _ = v63
            } else {
                var v64 int = 127
                switch ((v61 + v61) + v61) {
                case "\n\t":
                    if true {
                        if true {
                            fmt.Println(v59)
                            fmt.Println(v2, v2)
                            _ = len(v5)
                        }
                    }
                case "hello":
                    v64 = 42
                }
                _ = v64
            }
        }
        v5 = append(v5, 0)
        for fi65 := 0; fi65 < 1; fi65++ {
            delete(v4, "a")
            _ = fi65
        }
        _ = fi62
    }
    if (2147483647 >= 127) {
        var v66 uint64 = v59
        switch 128 {
        case -1:
            v6 -= v6
        case 4294967295:
            v1[v61] = ""
        case 32768:
            v60 = (^(v60))
        }
        var v67 uint64 = v66
        _ = v66
        _ = v67
    }
    v68 := GenS12{F0: v2, F1: 2}
    if (!(695 >= 4294967295)) {
        v2 = 65535
        if (("foo bar" <= "a") && (0 != (^(v6)))) {
            if ((v6 ^ (v6 - (v6 & (v6 - (v6 / (((-(v6))) | 1)))))) < (v6 / ((v6) | 1))) {
                if true {
                    fmt.Println(v60)
                    if false {
                        var v69 int16 = -654
                        var v70 int32 = (-(v6))
                        for fi71 := 0; fi71 < 8; fi71++ {
                            _ = len(v1)
                            v61 = v61
                            _ = fi71
                        }
                        _ = v69
                        _ = v70
                    }
                    v72 := GenS13{F0: (v60 ^ v60)}
                    _ = v72
                } else {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                }
                v60 -= (v60 - (v60 << ((255) & 63)))
                v3 += v3
            }
            fmt.Println(v2, v61)
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v59
    _ = v6
    _ = v60
    _ = v61
    _ = v68
    _ = v7
    _ = v8
}
