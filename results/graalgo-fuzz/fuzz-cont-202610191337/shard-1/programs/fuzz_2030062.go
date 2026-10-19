package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64; F1 uint32 }
type GenS3 struct { F0 uint16; F1 int8 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 string; F1 int64; F2 uint8 }
type GenS6 struct { F0 uint; F1 bool; F2 uint16 }
type GenS7 struct { F0 uint; F1 int16 }

func main() {
    var v1 uint = 42
    v2 := make(map[string]int)
    var v3 uint8 = 255
    v4 := []int{32767}
    var v5 int32 = 42
    if ("\n\t" > "") {
        var v6 int16 = int16(v3)
        var v7 uint8 = v3
        fmt.Println(v7, v3, v3)
        _ = v6
        _ = v7
    }
    switch "hello" {
    case "\n\t":
        fmt.Println(v3)
    default:
        if false {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            if (true || false) {
                for fi8 := 0; fi8 < 8; fi8++ {
                    if false {
                        for fi9 := 0; fi9 < 9; fi9++ {
                            var v10 uint = (v1 + (^(v1)))
                            v10 = (v1 * (v1 | (v1 + 441)))
                            fmt.Println(v5, v3, v10)
                            fmt.Println(v10)
                            _ = fi9
                            _ = v10
                        }
                        v5 = (^(v5))
                    }
                    if false {
                        v11 := GenS1{F0: uint16(v1)}
                        v3 += (v3 + 128)
                        var v12 int = fi8
                        _ = v11
                        _ = v12
                    }
                    if true {
                        var v13 float32 = (0.0 + -2.25)
                        _ = v13
                    }
                    _ = fi8
                }
            }
            var v14 uint = (^(v1))
            _ = v14
        } else {
            switch "hello" {
            case "\n\t":
                var v15 float64 = 1e10
                _ = v15
            case "a":
                v2["\n\t"] = 4294967295
            case "hello":
                if false {
                    v16 := GenS2{F0: 42, F1: 2147483647}
                    fmt.Println(v3, v3)
                    v3 += v3
                    _ = v16
                } else {
                    v5 += (^(v5))
                    v17 := GenS3{F0: 255, F1: -128}
                    if (-541 != 42) {
                        switch 32767 {
                        case 4294967295:
                            v4 = append(v4, 745)
                        case 9223372036854775807:
                            delete(v2, (((("a" + "a") + "") + "foo bar") + "\n\t"))
                        case -218:
                            switch 9223372036854775807 {
                            case 42:
                                _ = len(v4)
                            case 237:
                                v18 := GenS1{F0: 255}
                                _ = v18
                            case 2147483647:
                                v3 = v3
                            default:
                                var v19 uint32 = 42
                                _ = v19
                            }
                        default:
                            var v20 int16 = -762
                            _ = v20
                        }
                        fmt.Println(v5, v1)
                    }
                    _ = v17
                }
            default:
                for fi21 := 0; fi21 < 8; fi21++ {
                    v3 = v3
                    v22 := v2[""]
                    if (1 > 0) {
                        v3 -= 0
                        var v23 bool = (!(127 >= -1))
                        v24 := GenS2{F0: 2, F1: 42}
                        _ = v23
                        _ = v24
                    } else {
                        v25 := GenS3{F0: 327, F1: -26}
                        if true {
                            v26 := GenS4{F0: 42}
                            _ = v26
                        }
                        _ = v25
                    }
                    _ = fi21
                    _ = v22
                }
            }
            v27 := v2[(("a" + "foo bar") + "a")]
            _ = v27
        }
    }
    switch "hello" {
    case "foo bar":
        var v28 int16 = 959
        _ = v28
    case "\n\t":
        var v29 uint32 = 649
        _ = v29
    case "":
        var v30 int32 = v5
        _ = v30
    default:
        fmt.Println(v5)
    }
    v31 := GenS5{F0: "\n\t", F1: -9223372036854775808, F2: (v3 & (v3 / (((v3 ^ (^(v3)))) | 1)))}
    for fi32 := 0; fi32 < 14; fi32++ {
        v33 := GenS1{F0: 255}
        _ = fi32
        _ = v33
    }
    if len(v4) > 0 {
        v4[0] = -441
    }
    if (753 >= 0) {
        fmt.Println(v5)
    } else {
        for fi34 := 0; fi34 < 3; fi34++ {
            if true {
                switch fi34 {
                case -577:
                    v35 := GenS6{F0: 18446744073709551615, F1: (((1.5 >= 0.1) && ((fi34 & fi34) >= (^(fi34)))) && true), F2: 32767}
                    _ = v35
                case 42:
                    var v36 int = -269
                    _ = v36
                default:
                    if ((false || (false && (791 < 4294967295))) || ((255 != 2) && ((v1 + (^(v1))) > (^(v1))))) {
                        delete(v2, "hello")
                    } else {
                        v37 := GenS7{F0: 2, F1: 26}
                        _ = v37
                    }
                }
                for fi38 := 0; fi38 < 12; fi38++ {
                    v1 = (v1 | uint(v3))
                    v4 = append(v4, fi34)
                    _ = fi38
                }
                fmt.Println(v5, fi34)
            } else {
                v3 = v3
            }
            var v39 int32 = v5
            _ = fi34
            _ = v39
        }
    }
    v3 = (v3 << ((v1) & 63))
    if (!(-417 <= 127)) {
        v1 += v1
        _ = len(v2)
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    v2[("foo bar" + "foo bar")] = -672
    var v40 uint32 = 30
    v3 -= (v3 | 255)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v40
    _ = v5
}
