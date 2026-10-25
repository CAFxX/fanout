package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 uint8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 string; F1 uint8; F2 int64 }
type GenS4 struct { F0 int32; F1 int64; F2 int32 }
type GenS5 struct { F0 uint64; F1 uint16; F2 string }
type GenS6 struct { F0 bool; F1 int }
type GenS7 struct { F0 string }

func main() {
    var v1 uint = 2
    v2 := []uint32{}
    var v3 int32 = 42
    v4 := make(map[int]bool)
    var v5 int = 4294967295
    var v6 bool = false
    fmt.Println(v6, v5, v1)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v7 := GenS1{F0: int16(v3), F1: (-(v5)), F2: 2}
    fmt.Println(v1, v5)
    if (42 <= 262) {
        for fi8 := 0; fi8 < 16; fi8++ {
            v1 = 4294967295
            _ = fi8
        }
        switch ("hello" + "hello") {
        case "":
            _ = len(v4)
        case "foo bar":
            v9 := GenS1{F0: -279, F1: (v5 + (v5 << ((v1) & 63))), F2: 255}
            _ = v9
        default:
            v10 := GenS1{F0: 1000, F1: v5, F2: 255}
            _ = v10
        }
    } else {
        delete(v4, (v5 % ((1) | 1)))
    }
    var v11 int64 = 253
    v12 := GenS1{F0: 757, F1: v5, F2: 2}
    v13 := GenS2{F0: v1}
    v14 := v4[v5]
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println(v1)
    if true {
        var v15 int32 = v3
        if true {
            switch "foo bar" {
            case "":
                v16 := v4[(v5 ^ v5)]
                _ = v16
            default:
                for fi17 := 0; fi17 < 20; fi17++ {
                    if (v11 <= -9223372036854775808) {
                        v18 := GenS3{F0: "foo bar", F1: 1, F2: (v11 << ((v1) & 63))}
                        var v19 int64 = v11
                        _ = v18
                        _ = v19
                    }
                    if ((v1 ^ v1) != 2) {
                        for fi20 := 0; fi20 < 3; fi20++ {
                            var v21 uint32 = 102
                            fmt.Println(v6, v6)
                            v22 := GenS4{F0: int32(v21), F1: (v11 | v11), F2: v3}
                            _ = fi20
                            _ = v21
                            _ = v22
                        }
                    } else {
                        var v23 int64 = (v11 % ((v11) | 1))
                        _ = v23
                    }
                    if v6 {
                        if ((v1 >= (v1 & (v1 % ((v1) | 1)))) || (!(-1 != -1))) {
                            v5 = (-(v5))
                        } else {
                            var v24 int8 = -1
                            _ = v24
                        }
                        v25 := GenS3{F0: ("a" + (((("\n\t" + "a") + "a") + "") + "")), F1: 2, F2: 42}
                        _ = v25
                    }
                    _ = fi17
                }
            }
            if (42 == v5) {
                fmt.Println(v14, v6)
            } else {
                if v6 {
                    var v26 int = (v5 & (v5 % (((v5 & (v5 - v5))) | 1)))
                    v27 := GenS5{F0: 32, F1: 255, F2: "hello"}
                    fmt.Println(v11, v6, v26)
                    _ = v26
                    _ = v27
                } else {
                    var v28 float64 = float64(0.0)
                    v29 := GenS5{F0: 978, F1: 406, F2: "hello"}
                    for fi30 := 0; fi30 < 15; fi30++ {
                        for fi31 := 0; fi31 < 3; fi31++ {
                            switch "\n\t" {
                            case "":
                                fmt.Println(fi31, v3, v14)
                            default:
                                if (v6 && v14) {
                                    var v32 int64 = 218
                                    switch v5 {
                                    case 235:
                                        switch "\n\t" {
                                        case "\n\t":
                                            fmt.Println(v3)
                                        default:
                                            var v33 bool = (((v6 && (2 == 42)) || false) || (-553 == (fi30 << (((v1 ^ 1)) & 63))))
                                            _ = v33
                                        }
                                    case 0:
                                        fmt.Println(v32, v3, v32)
                                    case 2147483647:
                                        v1 = 1
                                    default:
                                        var v34 int8 = -128
                                        _ = v34
                                    }
                                    var v35 float64 = 0.1
                                    _ = v32
                                    _ = v35
                                } else {
                                    fmt.Println(fi30, v11, v5)
                                }
                            }
                            _ = fi31
                        }
                        if true {
                            var v36 bool = false
                            if (0 != 1) {
                                if v14 {
                                    v1 = 42
                                }
                                var v37 string = ""
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                _ = v37
                            } else {
                                v38 := GenS6{F0: (false && ((770 <= 42) && (v14 && v14))), F1: (fi30 << ((v1) & 63))}
                                _ = v38
                            }
                            _ = v36
                        } else {
                            var v39 uint32 = 2147483647
                            var v40 int = (v5 >> ((v1) & 63))
                            v41 := GenS4{F0: (^(v15)), F1: (v11 / ((v11) | 1)), F2: v15}
                            _ = v39
                            _ = v40
                            _ = v41
                        }
                        _ = fi30
                    }
                    _ = v28
                    _ = v29
                }
            }
            v4[(v5 + 4294967295)] = (128 != 103)
        } else {
            if (170 > 87) {
                v42 := GenS5{F0: 380, F1: 42, F2: "foo bar"}
                var v43 int32 = (v3 & v3)
                _ = v42
                _ = v43
            }
            v6 = (!v14)
            fmt.Println(v11, v6, v5)
        }
        _ = v15
    }
    v14 = v14
    switch v5 {
    case 4294967295:
        if (!(!v6)) {
            v3 -= 1
            for fi44 := 0; fi44 < 5; fi44++ {
                var v45 uint16 = 32767
                _ = fi44
                _ = v45
            }
        } else {
            switch "" {
            case "":
                for fi46 := 0; fi46 < 15; fi46++ {
                    switch v5 {
                    case 127:
                        var v47 uint = (v1 - v1)
                        _ = v47
                    case 716:
                        var v48 uint32 = 800
                        _ = v48
                    case 1:
                        v2 = append(v2, 626)
                    default:
                        var v49 uint16 = 132
                        _ = v49
                    }
                    if (v14 || ((-(v3)) <= v3)) {
                        var v50 uint8 = 255
                        _ = v50
                    }
                    _ = fi46
                }
            case "hello":
                switch ("\n\t" + (("a" + "a") + "foo bar")) {
                case "a":
                    if (1 > 65535) {
                        if (-32768 != -1) {
                            var v51 uint8 = 127
                            fmt.Println(v14, v3)
                            _ = v51
                        }
                        v52 := GenS7{F0: "foo bar"}
                        _ = v52
                    } else {
                        fmt.Println(v3, v14, v11)
                    }
                default:
                    if v6 {
                        v53 := GenS4{F0: v3, F1: 42, F2: v3}
                        switch "" {
                        case "\n\t":
                            switch (v5 * v5) {
                            case -2147483648:
                                v1 = v1
                            }
                        }
                        v11 -= v11
                        _ = v53
                    }
                }
            }
            _ = len(v2)
        }
    case -771:
        for fi54 := 0; fi54 < 15; fi54++ {
            v14 = ((v11 ^ -1) != (v11 << ((v1) & 63)))
            _ = fi54
        }
    case 241:
        v6 = (0 <= 42)
    default:
        for fi55 := 0; fi55 < 15; fi55++ {
            var v56 float32 = 1e10
            _ = fi55
            _ = v56
        }
    }
    var v57 uint8 = 255
    v14 = v14
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v57:", v57)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v57
    _ = v6
    _ = v7
}
