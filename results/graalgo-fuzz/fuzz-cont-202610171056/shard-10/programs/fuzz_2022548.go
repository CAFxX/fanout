package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int; F1 bool }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int16; F1 uint16; F2 uint8 }
func genFn0() int64 {
    var v4 uint64 = 18446744073709551615
    v4 = (v4 << ((65535) & 63))
    fmt.Println(v4)
    var v5 string = "foo bar"
    _ = v4
    _ = v5
    return -9223372036854775808
}

func main() {
    var v1 uint32 = 2147483647
    v2 := make(map[string]int)
    var v3 float32 = (-2.25 - 3.14159)
    v6 := []int{2147483647, 42}
    v1 += v1
    v1 = (^(v1))
    v2["hello"] = 128
    v3 = 0.0
    v7 := GenS1{F0: -2147483648}
    _ = len(v6)
    switch 1 {
    case -1:
        for fi8 := 0; fi8 < 7; fi8++ {
            v1 = (v1 - 250)
            if (0 != 42) {
                if (!(false || (32767 <= 1000))) {
                    for fi9 := 0; fi9 < 1; fi9++ {
                        fmt.Println(fi9)
                        _ = fi9
                    }
                    switch "a" {
                    case "a":
                        v3 -= v3
                    default:
                        v10 := GenS2{F0: 128, F1: false}
                        _ = v10
                    }
                }
                for fi11 := 0; fi11 < 16; fi11++ {
                    v1 += 1
                    _ = fi11
                }
            }
            for fi12 := 0; fi12 < 13; fi12++ {
                v2["hello"] = (fi8 | fi12)
                v3 -= v3
                if true {
                    v2["\n\t"] = fi8
                }
                _ = fi12
            }
            _ = fi8
        }
    default:
        v13 := GenS2{F0: 961, F1: true}
        _ = v13
    }
    if (!(!false)) {
        var v14 uint16 = 0
        _ = v14
    } else {
        var v15 int16 = 0
        var v16 bool = (-2147483648 < -739)
        if (4294967295 > 519) {
            v17 := GenS1{F0: -383}
            if false {
                if (1e10 != 3.14159) {
                    if ("foo bar" >= "") {
                        fmt.Println(v1, v1, v15)
                        fmt.Println(v15, v16)
                    }
                    for fi18 := 0; fi18 < 7; fi18++ {
                        var v19 uint8 = 255
                        fmt.Println(v3, fi18)
                        for fi20 := 0; fi20 < 18; fi20++ {
                            v16 = (!v16)
                            var v21 uint32 = 469
                            v22 := GenS2{F0: 4294967295, F1: v16}
                            _ = fi20
                            _ = v21
                            _ = v22
                        }
                        _ = fi18
                        _ = v19
                    }
                } else {
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    switch 127 {
                    case 1:
                        fmt.Println(v16)
                    case 128:
                        var v23 uint32 = v1
                        _ = v23
                    case 9223372036854775807:
                        v24 := GenS1{F0: 4294967295}
                        _ = v24
                    default:
                        v25 := GenS3{F0: 0}
                        _ = v25
                    }
                    fmt.Println(v1)
                }
                var v26 int8 = -52
                switch "\n\t" {
                case "a":
                    if (v16 || false) {
                        var v27 float64 = (float64(((1e10 - 3.14159) * v3)) * 0.1)
                        v27 += v27
                        _ = v27
                    }
                case "foo bar":
                    v28 := v2[""]
                    _ = v28
                case "\n\t":
                    for fi29 := 0; fi29 < 16; fi29++ {
                        var v30 uint8 = 127
                        v3 = -2.25
                        _ = fi29
                        _ = v30
                    }
                default:
                    if (!((-2147483648 < -2147483648) && (v16 && ((^(v15)) >= (v15 + (^(v15))))))) {
                        v15 += (v15 << ((65535) & 63))
                    } else {
                        if (-2147483648 > -1) {
                            if (v16 && ((0.0 < 1.5) && false)) {
                                v1 += 246
                                v26 = v26
                            } else {
                                if v16 {
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                    v3 += v3
                                    for fi31 := 0; fi31 < 1; fi31++ {
                                        var v32 uint32 = (v1 + v1)
                                        if ((v15 | v15) == (v15 ^ (v15 - (v15 - (v15 + v15))))) {
                                            for fi33 := 0; fi33 < 6; fi33++ {
                                                v34 := GenS2{F0: 127, F1: (519 < 42)}
                                                v15 += (v15 << ((1) & 63))
                                                _ = fi33
                                                _ = v34
                                            }
                                        }
                                        var v35 int8 = v26
                                        _ = fi31
                                        _ = v32
                                        _ = v35
                                    }
                                } else {
                                    v36 := GenS1{F0: 0}
                                    v37 := GenS4{F0: (v15 & (v15 << ((1) & 63))), F1: 42, F2: 1}
                                    switch "" {
                                    case "hello":
                                        v26 -= (-(v26))
                                    }
                                    _ = v36
                                    _ = v37
                                }
                                for fi38 := 0; fi38 < 5; fi38++ {
                                    v16 = ((v15 + 0) <= (v15 << ((18446744073709551615) & 63)))
                                    fmt.Println(fi38, v3)
                                    v1 = v1
                                    _ = fi38
                                }
                                var v39 int16 = v15
                                _ = v39
                            }
                            v1 = (^(v1))
                            for fi40 := 0; fi40 < 8; fi40++ {
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                                _ = fi40
                            }
                        }
                        var v41 int16 = v15
                        _ = v41
                    }
                }
                _ = v26
            } else {
                _ = len(v2)
                if (4294967295 < v1) {
                    fmt.Println(v15, v16, v3)
                    v16 = true
                } else {
                    for fi42 := 0; fi42 < 7; fi42++ {
                        v43 := GenS4{F0: v15, F1: 65535, F2: 0}
                        _ = fi42
                        _ = v43
                    }
                }
            }
            _ = v17
        }
        _ = v15
        _ = v16
    }
    fmt.Println(v3)
    if (!true) {
        var v44 int32 = 42
        for fi45 := 0; fi45 < 20; fi45++ {
            fmt.Println(v1, v44, fi45)
            _ = fi45
        }
        v44 -= (-(v44))
        _ = v44
    } else {
        if (255 < 127) {
            var v46 uint64 = 734
            v3 -= float32((0.1 / float64(v3)))
            _ = v46
        } else {
            v3 = ((v3 - v3) / 0.1)
            v47 := v2[""]
            v47 -= v47
            _ = v47
        }
    }
    fmt.Println(v3)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v6
    _ = v7
}
