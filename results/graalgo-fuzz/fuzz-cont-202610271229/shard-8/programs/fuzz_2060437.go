package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 string; F1 int32 }
type GenS3 struct { F0 uint; F1 uint16; F2 bool }
type GenS4 struct { F0 int8; F1 int }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint32; F1 int32; F2 uint8 }
type GenS8 struct { F0 bool; F1 int32 }
type GenS9 struct { F0 uint8; F1 uint32 }
type GenS10 struct { F0 uint; F1 bool; F2 uint64 }
type GenS11 struct { F0 int64; F1 uint32 }
type GenS12 struct { F0 string }
type GenS13 struct { F0 int16 }
type GenS14 struct { F0 string; F1 uint64 }

func main() {
    v1 := make(map[string]string)
    v2 := []string{("foo bar" + "hello"), "foo bar"}
    var v3 string = ("\n\t" + "\n\t")
    v4 := make(map[string]bool)
    var v5 int8 = -127
    fmt.Println(v3, v5)
    for fi6 := 0; fi6 < 14; fi6++ {
        var v7 uint8 = 127
        v8 := GenS1{F0: 404, F1: (v7 + 71)}
        switch 969 {
        case 521:
            var v9 uint64 = 1
            _ = v9
        default:
            for fi10 := 0; fi10 < 1; fi10++ {
                switch 0 {
                case 32767:
                    v11 := GenS2{F0: ("foo bar" + v3), F1: -2147483648}
                    _ = v11
                case -9223372036854775808:
                    v5 = v5
                default:
                    fmt.Println(v3, v7, fi10)
                }
                var v12 float64 = 1.5
                v3 = ("\n\t" + v3)
                _ = fi10
                _ = v12
            }
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    switch -112 {
    case -473:
        v4[(("foo bar" + (("hello" + "") + v3)) + "")] = true
    default:
        for fi13 := 0; fi13 < 2; fi13++ {
            fmt.Println(v5, v3, v5)
            var v14 uint8 = 93
            _ = fi13
            _ = v14
        }
    }
    if ((255 <= 0) || true) {
        v5 += v5
    }
    v15 := GenS1{F0: 255, F1: 128}
    var v16 uint = 338
    for fi17 := 0; fi17 < 15; fi17++ {
        _ = len(v1)
        _ = len(v4)
        _ = fi17
    }
    fmt.Println(v5, v3)
    v18 := GenS3{F0: v16, F1: 2, F2: false}
    if true {
        var v19 int16 = -32768
        var v20 uint = v16
        _ = v19
        _ = v20
    } else {
        if (0 >= 0) {
            v21 := GenS4{F0: 100, F1: 139}
            switch "hello" {
            case "\n\t":
                var v22 int8 = (v5 + v5)
                _ = v22
            default:
                fmt.Println(v3, v16)
            }
            _ = v21
        }
        v23 := GenS5{F0: (v16 / ((v16) | 1))}
        switch "" {
        case "foo bar":
            if (!true) {
                fmt.Println(v5, v3)
                switch v3 {
                case "hello":
                    v24 := GenS6{F0: 2}
                    _ = v24
                case "a":
                    var v25 uint8 = 221
                    _ = v25
                case "":
                    v5 -= (v5 | v5)
                default:
                    var v26 float64 = (3.14159 - 1.5)
                    _ = v26
                }
            } else {
                switch ("\n\t" + v3) {
                case "hello":
                    for fi27 := 0; fi27 < 9; fi27++ {
                        fmt.Println(v5)
                        v2 = append(v2, v3)
                        var v28 uint8 = 253
                        _ = fi27
                        _ = v28
                    }
                case "\n\t":
                    if (1 > 598) {
                        var v29 int64 = 42
                        _ = v29
                    } else {
                        var v30 int8 = (^(v5))
                        for fi31 := 0; fi31 < 5; fi31++ {
                            _ = len(v2)
                            if false {
                                v16 = v16
                                var v32 int8 = (v5 / ((127) | 1))
                                v33 := GenS5{F0: v16}
                                _ = v32
                                _ = v33
                            } else {
                                for fi34 := 0; fi34 < 6; fi34++ {
                                    if false {
                                        switch 32768 {
                                        case 9223372036854775807:
                                            var v35 uint = (^(v16))
                                            _ = v35
                                        default:
                                            var v36 bool = true
                                            _ = v36
                                        }
                                        fmt.Println(fi31, v16, fi31)
                                    } else {
                                        v37 := GenS6{F0: 831}
                                        v38 := GenS7{F0: 2, F1: 1, F2: 32}
                                        if true {
                                            _ = len(v2)
                                        } else {
                                            v39 := GenS8{F0: (!false), F1: -2147483648}
                                            if true {
                                                _ = len(v2)
                                                fmt.Println(v3, fi34, fi34)
                                            }
                                            _ = v39
                                        }
                                        v40 := GenS9{F0: uint8(v16), F1: uint32(v16)}
                                        fmt.Println(fi34, v5)
                                        _ = v37
                                        _ = v38
                                        _ = v40
                                    }
                                    var v41 int32 = -490
                                    v3 = "hello"
                                    _ = fi34
                                    _ = v41
                                }
                            }
                            v42 := GenS2{F0: v3, F1: 866}
                            _ = fi31
                            _ = v42
                        }
                        var v43 float32 = ((0.0 - 0.1) + 1e10)
                        _ = v30
                        _ = v43
                    }
                case "":
                    v44 := v4[((v3 + ((v3 + (v3 + ((v3 + v3) + v3))) + "foo bar")) + "")]
                    _ = v44
                default:
                    switch 1 {
                    case 4294967295:
                        var v45 uint8 = 2
                        _ = v45
                    case -745:
                        for fi46 := 0; fi46 < 12; fi46++ {
                            if (255 <= 1) {
                                fmt.Println(v3, v16, v16)
                            }
                            v47 := GenS10{F0: 448, F1: false, F2: 625}
                            _ = fi46
                            _ = v47
                        }
                    }
                }
                var v48 uint64 = 4294967295
                _ = v48
            }
        case "\n\t":
            for fi49 := 0; fi49 < 13; fi49++ {
                for fi50 := 0; fi50 < 18; fi50++ {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = fi50
                }
                v5 += -1
                for fi51 := 0; fi51 < 3; fi51++ {
                    if (!true) {
                        var v52 int16 = 32767
                        v3 = (v3 + v3)
                        _ = v52
                    } else {
                        var v53 uint64 = 550
                        fmt.Println(v3, v3)
                        v54 := GenS11{F0: 42, F1: 928}
                        _ = v53
                        _ = v54
                    }
                    v16 -= v16
                    var v55 bool = true
                    _ = fi51
                    _ = v55
                }
                _ = fi49
            }
        case "":
            if (v5 > 1) {
                var v56 int = 32767
                for fi57 := 0; fi57 < 6; fi57++ {
                    var v58 float32 = ((1e10 + float32(-2.25)) - 0.0)
                    _ = fi57
                    _ = v58
                }
                _ = v56
            } else {
                for fi59 := 0; fi59 < 15; fi59++ {
                    if true {
                        var v60 uint16 = 255
                        var v61 uint = 255
                        var v62 uint64 = 389
                        _ = v60
                        _ = v61
                        _ = v62
                    }
                    v16 -= v16
                    _ = fi59
                }
                for fi63 := 0; fi63 < 12; fi63++ {
                    var v64 bool = true
                    for fi65 := 0; fi65 < 20; fi65++ {
                        var v66 uint64 = 12
                        v67 := GenS3{F0: (v16 % ((827) | 1)), F1: 1, F2: (v66 <= (v66 & v66))}
                        if false {
                            var v68 int32 = 1
                            v64 = (!(v64 && (319 >= 65535)))
                            _ = v68
                        }
                        _ = fi65
                        _ = v66
                        _ = v67
                    }
                    _ = fi63
                    _ = v64
                }
            }
        }
        _ = v23
    }
    v69 := GenS6{F0: 713}
    switch 128 {
    case -1:
        v70 := GenS12{F0: v3}
        _ = v70
    case 32768:
        v71 := GenS13{F0: -32768}
        _ = v71
    }
    var v72 int64 = 42
    v73 := GenS14{F0: v3, F1: 18446744073709551615}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v72 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v15
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v69
    _ = v72
    _ = v73
}
