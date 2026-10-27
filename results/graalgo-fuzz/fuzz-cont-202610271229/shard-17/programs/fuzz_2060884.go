package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint64; F1 int32; F2 string }
type GenS4 struct { F0 bool; F1 bool }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int64; F1 uint16; F2 int64 }
func genFn0(p0 int8, p1 uint32, p2 uint64) uint64 {
    var v4 uint64 = 42
    var v5 int32 = 614
    if true {
        for fi6 := 0; fi6 < 14; fi6++ {
            v7 := GenS1{F0: p1}
            for fi8 := 0; fi8 < 3; fi8++ {
                switch "foo bar" {
                case "a":
                    p1 += (p1 ^ p1)
                default:
                    var v9 int16 = -1
                    _ = v9
                }
                _ = fi8
            }
            _ = fi6
            _ = v7
        }
        v5 = v5
    } else {
        for fi10 := 0; fi10 < 2; fi10++ {
            var v11 int16 = 42
            if false {
                v5 = (^(v5))
            }
            _ = fi10
            _ = v11
        }
        if true {
            switch 119 {
            case 128:
                if (!((!(!true)) || false)) {
                    v12 := []int{2147483647, 0, 0, 2147483647}
                    var v13 float32 = float32((0.1 / 1e10))
                    if (!((0.0 - 3.14159) <= 1e10)) {
                        if true {
                            if (!false) {
                                fmt.Println(p1)
                            } else {
                                if true {
                                    var v14 uint = 2
                                    _ = v14
                                } else {
                                    v15 := GenS2{F0: -2147483648}
                                    _ = v15
                                }
                                switch 1 {
                                case -816:
                                    v16 := GenS3{F0: (v4 << ((0) & 63)), F1: (v5 << ((255) & 63)), F2: ""}
                                    _ = v16
                                default:
                                    var v17 int16 = 0
                                    _ = v17
                                }
                                p0 = 0
                            }
                        } else {
                            switch "a" {
                            case "a":
                                for fi18 := 0; fi18 < 10; fi18++ {
                                    if true {
                                        fmt.Println(v4, v4, v5)
                                        v19 := GenS4{F0: false, F1: (fi18 != (fi18 + (fi18 / (((fi18 / (((^(fi18))) | 1))) | 1))))}
                                        v13 = 0.0
                                        _ = v19
                                    } else {
                                        p1 -= p1
                                        if ((639 != 42) || false) {
                                            v20 := GenS2{F0: v5}
                                            v5 += 0
                                            _ = v20
                                        } else {
                                            p1 += (p1 % ((2) | 1))
                                            v12 = append(v12, fi18)
                                            _ = len(v12)
                                        }
                                    }
                                    p0 = 100
                                    _ = fi18
                                }
                            case "hello":
                                var v21 uint32 = (p1 ^ (p1 + p1))
                                _ = v21
                            case "\n\t":
                                if len(v12) > 0 {
                                    v12[0] = 127
                                }
                            default:
                                v12 = append(v12, 641)
                            }
                        }
                    } else {
                        if len(v12) >= 2 {
                            v12 = v12[1:len(v12)]
                        }
                        p0 += p0
                        if len(v12) > 0 {
                            v12[0] = 4294967295
                        }
                    }
                    _ = v12
                    _ = v13
                } else {
                    switch ("foo bar" + "") {
                    case "\n\t":
                        v22 := []int{32767, 2147483647}
                        _ = v22
                    case "":
                        if (((1.5 - ((1.5 + 0.0) / 1e10)) - float64(-2.25)) >= (0.0 + ((-2.25 / 1.0) * (3.14159 / (3.14159 * (0.0 - (-2.25 + (((1.5 + 1.5) - 0.1) * (0.1 * (1.5 / 1e10)))))))))) {
                            for fi23 := 0; fi23 < 15; fi23++ {
                                var v24 uint = 65535
                                p0 = (p0 * p0)
                                v25 := []int{(fi23 + -150), -2147483648}
                                _ = fi23
                                _ = v24
                                _ = v25
                            }
                            v5 = 42
                        }
                    }
                }
            case -439:
                v26 := []int{-118, 2147483647, 42, -165}
                _ = v26
            }
        }
        for fi27 := 0; fi27 < 4; fi27++ {
            v28 := GenS5{F0: -128}
            v29 := []int{(-(fi27))}
            _ = fi27
            _ = v28
            _ = v29
        }
    }
    p2 -= 42
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return (v4 ^ (p2 + (v4 * 18446744073709551615)))
}

func main() {
    v1 := []string{""}
    var v2 float64 = ((0.1 * (1e10 * ((0.0 / 1.5) - (float64(-2.25) - (float64(0.0) + 0.1))))) + (3.14159 / ((1.5 / 0.1) / (-2.25 * 3.14159))))
    v3 := []uint32{1, 342}
    if false {
        if len(v1) > 0 {
            v1[0] = "foo bar"
        }
    }
    v30 := GenS5{F0: 1}
    var v31 uint32 = 760
    fmt.Println(v2, v31)
    v2 = (0.0 / (v2 / ((v2 + v2) + v2)))
    v32 := GenS3{F0: genFn0(105, v31, 0), F1: -2147483648, F2: ("\n\t" + "")}
    if true {
        if false {
            for fi33 := 0; fi33 < 16; fi33++ {
                v2 = v2
                if ((int32(fi33) == -606) || (uint16(v31) != 32767)) {
                    var v34 bool = true
                    _ = v34
                } else {
                    for fi35 := 0; fi35 < 10; fi35++ {
                        switch "\n\t" {
                        case "hello":
                            v36 := GenS6{F0: 602, F1: 408, F2: 9223372036854775807}
                            _ = v36
                        case "\n\t":
                            if len(v3) > 0 {
                                v3[0] = (v31 ^ (v31 ^ v31))
                            }
                        default:
                            _ = len(v3)
                        }
                        _ = fi35
                    }
                    fmt.Println(v2, v31)
                    v31 -= (v31 | v31)
                }
                v31 = (v31 / ((1) | 1))
                _ = fi33
            }
            var v37 uint16 = 1
            if (!(255 > 2)) {
                if true {
                    v2 = v2
                    v38 := GenS4{F0: (!(0 == 65535)), F1: true}
                    if (!false) {
                        if (!((!false) && (genFn0(-1, (v31 & 4294967295), 18446744073709551615) <= 786))) {
                            v39 := GenS2{F0: 1}
                            fmt.Println(v37, v37, v37)
                            fmt.Println(v31, v2)
                            _ = v39
                        } else {
                            var v40 float64 = (v2 / ((v2 + v2) + v2))
                            _ = v40
                        }
                    } else {
                        v2 = v2
                        _ = len(v3)
                    }
                    _ = v38
                }
            } else {
                var v41 uint64 = 624
                _ = v41
            }
            _ = v37
        }
    }
    for fi42 := 0; fi42 < 10; fi42++ {
        var v43 uint32 = v31
        for fi44 := 0; fi44 < 14; fi44++ {
            var v45 uint = 4294967295
            _ = fi44
            _ = v45
        }
        fmt.Println(v31)
        _ = fi42
        _ = v43
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    switch 128 {
    case -2147483648:
        v2 -= 3.14159
    case -1:
        var v46 uint16 = 58
        _ = v46
    case 37:
        var v47 int = 42
        _ = v47
    }
    if len(v1) > 0 {
        v1[0] = "\n\t"
    }
    v31 -= 771
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
}
