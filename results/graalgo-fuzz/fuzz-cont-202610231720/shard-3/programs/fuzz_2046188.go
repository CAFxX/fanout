package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int32; F1 bool; F2 bool }
type GenS3 struct { F0 int; F1 uint16; F2 string }
type GenS4 struct { F0 uint8; F1 int; F2 uint8 }
func genFn0(p0 uint8, p1 int16, p2 uint16) int32 {
    for fi6 := 0; fi6 < 9; fi6++ {
        var v7 uint = 402
        var v8 int16 = (-(p1))
        _ = fi6
        _ = v7
        _ = v8
    }
    if false {
        p1 += (p1 * -465)
    }
    var v9 uint16 = 65535
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return -155
}

func main() {
    var v1 bool = (-1 <= -9223372036854775808)
    var v2 int32 = -2147483648
    v3 := make(map[int]int)
    var v4 float32 = 1e10
    var v5 int16 = -10
    v10 := GenS1{F0: 85}
    v5 += v5
    if v1 {
        v11 := []int{}
        _ = v11
    }
    v1 = (42 > 545)
    if v1 {
        v12 := GenS1{F0: 18446744073709551615}
        var v13 int = 128
        fmt.Println(v13, v13, v5)
        _ = v12
        _ = v13
    } else {
        v14 := []int{128, 32768, 9223372036854775807}
        _ = v14
    }
    if ((!(float32(((float64(1.5) + 3.14159) - -2.25)) != (v4 / v4))) || false) {
        for fi15 := 0; fi15 < 16; fi15++ {
            if ((624 <= 1) || v1) {
                v16 := []int{fi15, -2147483648}
                _ = v16
            } else {
                var v17 int8 = 127
                if false {
                    v5 -= (-(v5))
                } else {
                    v18 := []int{fi15}
                    v19 := GenS1{F0: 760}
                    for fi20 := 0; fi20 < 15; fi20++ {
                        var v21 int16 = (v5 | 469)
                        _ = fi20
                        _ = v21
                    }
                    _ = v18
                    _ = v19
                }
                for fi22 := 0; fi22 < 10; fi22++ {
                    switch "a" {
                    case "":
                        switch fi15 {
                        case 2147483647:
                            if v1 {
                                var v23 uint32 = 42
                                v2 = 1
                                v24 := v3[(fi22 & (fi22 / (((fi15 - (fi15 ^ fi15))) | 1)))]
                                _ = v23
                                _ = v24
                            } else {
                                if ((^(fi15)) > fi22) {
                                    var v25 int = int(v2)
                                    v26 := GenS2{F0: (v2 - (v2 * v2)), F1: (!v1), F2: (42 >= 18446744073709551615)}
                                    if true {
                                        fmt.Println(v4, v4)
                                    }
                                    switch "\n\t" {
                                    case "":
                                        switch "" {
                                        case "foo bar":
                                            if v1 {
                                                switch ("" + "") {
                                                case "":
                                                    fmt.Println(fi22, v5, fi22)
                                                }
                                            }
                                        case "hello":
                                            var v27 float64 = (float64(float32((0.1 - 0.1))) * (0.0 / -2.25))
                                            _ = v27
                                        default:
                                            var v28 int8 = v17
                                            _ = v28
                                        }
                                    case "hello":
                                        switch fi22 {
                                        case 32767:
                                            var v29 bool = v1
                                            _ = v29
                                        case 42:
                                            switch v25 {
                                            case -9223372036854775808:
                                                var v30 float64 = float64((1e10 / v4))
                                                v31 := GenS2{F0: genFn0(255, v5, 1), F1: true, F2: (!(v1 && v1))}
                                                _ = v30
                                                _ = v31
                                            }
                                        case 32768:
                                            v2 = (v2 - 2147483647)
                                        default:
                                            v32 := GenS1{F0: 18446744073709551615}
                                            _ = v32
                                        }
                                    case "\n\t":
                                        var v33 float64 = 1e10
                                        _ = v33
                                    default:
                                        if (32767 > 1) {
                                            v34 := GenS2{F0: (v2 % ((2147483647) | 1)), F1: true, F2: (!true)}
                                            v35 := GenS2{F0: v2, F1: (v1 && true), F2: (!v1)}
                                            _ = v34
                                            _ = v35
                                        }
                                    }
                                    _ = v25
                                    _ = v26
                                }
                            }
                            v1 = true
                            if (255 >= 127) {
                                var v36 float64 = 1.5
                                v37 := GenS3{F0: (fi15 | fi15), F1: 0, F2: "foo bar"}
                                v38 := GenS4{F0: 128, F1: 128, F2: 172}
                                _ = v36
                                _ = v37
                                _ = v38
                            }
                            fmt.Println(fi15)
                        case 32767:
                            v2 -= (v2 % ((v2) | 1))
                        case -40:
                            var v39 uint8 = 127
                            _ = v39
                        }
                    }
                    fmt.Println(v2, fi15, v2)
                    var v40 uint = 42
                    _ = fi22
                    _ = v40
                }
                _ = v17
            }
            var v41 uint32 = 0
            _ = fi15
            _ = v41
        }
    } else {
        var v42 uint = 65535
        _ = v42
    }
    fmt.Println(v5)
    for fi43 := 0; fi43 < 8; fi43++ {
        var v44 int64 = 885
        v3[fi43] = fi43
        _ = fi43
        _ = v44
    }
    v4 -= (1.5 * (float32(float64(float32(1e10))) - v4))
    var v45 int = 142
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v5
}
