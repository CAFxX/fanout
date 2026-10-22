package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 int32; F1 int32 }
type GenS3 struct { F0 uint; F1 uint16; F2 string }
type GenS4 struct { F0 uint16; F1 int64; F2 uint }
type GenS5 struct { F0 int; F1 uint8; F2 int16 }
type GenS6 struct { F0 string; F1 int; F2 uint16 }
type GenS7 struct { F0 int16; F1 int16 }
func genFn0(p0 int32) uint16 {
    v6 := GenS1{F0: 127, F1: p0}
    for fi7 := 0; fi7 < 2; fi7++ {
        v8 := GenS1{F0: 255, F1: 981}
        p0 = p0
        _ = fi7
        _ = v8
    }
    v9 := []int{0, 42, 127, -133}
    if false {
        p0 += 879
    }
    for fi10 := 0; fi10 < 4; fi10++ {
        for fi11 := 0; fi11 < 13; fi11++ {
            p0 += 42
            _ = fi11
        }
        if (false && (!true)) {
            fmt.Println(p0, p0)
            if len(v9) > 0 {
                v9[0] = (fi10 >> ((0) & 63))
            }
        } else {
            v12 := GenS2{F0: p0, F1: p0}
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
            _ = v12
        }
        if (4294967295 == 1) {
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
            if (float64(float32(0.1)) >= 0.0) {
                if false {
                    for fi13 := 0; fi13 < 15; fi13++ {
                        switch fi10 {
                        case 1:
                            if len(v9) >= 2 {
                                v9 = v9[1:len(v9)]
                            }
                        case 449:
                            p0 += 1
                        case 127:
                            switch "foo bar" {
                            case "foo bar":
                                p0 = p0
                            case "a":
                                p0 -= p0
                            case "hello":
                                var v14 uint32 = 335
                                _ = v14
                            }
                        default:
                            if (0.1 > 1e10) {
                                if len(v9) > 0 {
                                    v9[0] = fi13
                                }
                            }
                            p0 = (p0 + -2147483648)
                        }
                        if true {
                            v15 := GenS1{F0: 127, F1: p0}
                            _ = v15
                        } else {
                            switch ("\n\t" + "hello") {
                            case "a":
                                var v16 int = (fi10 + (fi13 & fi10))
                                _ = v16
                            case "foo bar":
                                v17 := GenS1{F0: 2, F1: (p0 % ((p0) | 1))}
                                _ = v17
                            default:
                                if false {
                                    if (234 > fi10) {
                                        var v18 float64 = 1e10
                                        _ = v18
                                    }
                                    _ = len(v9)
                                    fmt.Println(fi13, fi10)
                                    if true {
                                        p0 += p0
                                        if (327 != -9223372036854775808) {
                                            switch fi10 {
                                            case 32767:
                                                p0 = 42
                                            case 4294967295:
                                                v19 := GenS2{F0: p0, F1: (p0 & (p0 / ((1) | 1)))}
                                                _ = v19
                                            default:
                                                _ = len(v9)
                                            }
                                            if (!(p0 != (p0 ^ (p0 - (p0 + (p0 % ((p0) | 1))))))) {
                                                p0 -= p0
                                                var v20 uint32 = 954
                                                if true {
                                                    var v21 uint8 = 128
                                                    _ = v21
                                                }
                                                _ = v20
                                            } else {
                                                p0 -= p0
                                                switch "foo bar" {
                                                case "\n\t":
                                                    p0 = p0
                                                case "foo bar":
                                                    _ = len(v9)
                                                }
                                                v22 := GenS3{F0: 4294967295, F1: 65535, F2: "a"}
                                                if (((((0.0 * 3.14159) - (1.5 / 0.1)) - 0.1) + 1.5) >= 1e10) {
                                                    if len(v9) >= 2 {
                                                        v9 = v9[1:len(v9)]
                                                    }
                                                    var v23 uint16 = 2
                                                    p0 += p0
                                                    _ = v23
                                                } else {
                                                    switch fi13 {
                                                    case -2147483648:
                                                        p0 -= p0
                                                    }
                                                    p0 -= (p0 / (((^(p0))) | 1))
                                                    v24 := GenS1{F0: 127, F1: p0}
                                                    _ = v24
                                                }
                                                _ = v22
                                            }
                                            p0 = 42
                                        }
                                        var v25 uint8 = 75
                                        _ = v25
                                    }
                                } else {
                                    var v26 float32 = -2.25
                                    switch (fi13 >> ((1) & 63)) {
                                    case -868:
                                        v26 = v26
                                    default:
                                        fmt.Println(fi10, fi10)
                                    }
                                    if (("foo bar" <= (("" + "foo bar") + "\n\t")) && true) {
                                        fmt.Println(v26, p0, v26)
                                        switch 0 {
                                        case 163:
                                            v9 = append(v9, (fi10 >> ((65535) & 63)))
                                        case 449:
                                            var v27 int8 = 79
                                            _ = v27
                                        }
                                        if (false || false) {
                                            p0 = p0
                                            var v28 float32 = v26
                                            _ = v28
                                        }
                                    } else {
                                        v26 -= v26
                                        v29 := GenS4{F0: 32767, F1: 248, F2: 255}
                                        _ = v29
                                    }
                                    v26 = v26
                                    fmt.Println(v26, fi10)
                                    _ = v26
                                }
                            }
                        }
                        _ = fi13
                    }
                    var v30 uint32 = 1
                    for fi31 := 0; fi31 < 1; fi31++ {
                        v32 := GenS5{F0: (fi31 - 4294967295), F1: 13, F2: -32768}
                        var v33 int8 = 42
                        _ = fi31
                        _ = v32
                        _ = v33
                    }
                    _ = v30
                }
                for fi34 := 0; fi34 < 17; fi34++ {
                    fmt.Println(p0, p0, p0)
                    _ = fi34
                }
                if true {
                    v35 := GenS6{F0: "\n\t", F1: 128, F2: 42}
                    v36 := GenS6{F0: "hello", F1: int(p0), F2: 255}
                    _ = v35
                    _ = v36
                } else {
                    var v37 int8 = 122
                    _ = v37
                }
            }
        }
        _ = fi10
    }
    _ = p0
    _ = v6
    _ = v9
    return 2
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint = 18446744073709551615
    v3 := []int{-970, -9223372036854775808}
    v4 := make(map[uint32]int)
    v5 := []string{}
    v38 := GenS4{F0: 32767, F1: 1, F2: v2}
    v39 := GenS7{F0: 1, F1: 923}
    v2 -= v2
    for fi40 := 0; fi40 < 10; fi40++ {
        var v41 float64 = float64(1e10)
        for fi42 := 0; fi42 < 9; fi42++ {
            var v43 uint16 = genFn0(2147483647)
            var v44 float64 = (v41 / ((((0.1 - ((v41 / v41) * ((v41 + 3.14159) / ((v41 - float64((0.1 * 0.0))) - 0.0)))) + (0.1 * v41)) / float64((1.5 / (3.14159 / 1e10)))) / v41))
            _ = fi42
            _ = v43
            _ = v44
        }
        switch (fi40 * 414) {
        case 926:
            if ((fi40 << ((519) & 63)) == fi40) {
                var v45 float64 = v41
                _ = v45
            }
        }
        _ = fi40
        _ = v41
    }
    var v46 int = 42
    if (int32(v2) <= -1) {
        for fi47 := 0; fi47 < 3; fi47++ {
            fmt.Println(v46, v46)
            switch fi47 {
            case 896:
                v2 += (^(v2))
            }
            _ = len(v4)
            _ = fi47
        }
        var v48 bool = (!true)
        _ = v48
    }
    v46 -= (v46 * v46)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v2 = (^(v2))
    for fi49 := 0; fi49 < 17; fi49++ {
        var v50 int8 = 127
        _ = fi49
        _ = v50
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v46
    _ = v5
}
