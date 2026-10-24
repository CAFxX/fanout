package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string; F1 string; F2 uint8 }
type GenS3 struct { F0 int64; F1 uint32 }
type GenS4 struct { F0 int }

func main() {
    var v1 uint32 = 928
    var v2 uint64 = uint64(v1)
    var v3 float64 = ((1e10 + float64(0.0)) + 0.0)
    v4 := []int{42, 0, 4294967295, -1}
    v5 := make(map[uint32]bool)
    var v6 int16 = 42
    switch 32767 {
    case 4294967295:
        var v7 int8 = 127
        _ = v7
    case -211:
        v8 := GenS1{F0: v1}
        _ = v8
    default:
        var v9 int32 = -1
        _ = v9
    }
    v6 += (-(v6))
    var v10 float32 = (-2.25 + 1.5)
    v2 += 42
    fmt.Println(v3, v10, v1)
    fmt.Println(v1, v1)
    for fi11 := 0; fi11 < 5; fi11++ {
        if false {
            for fi12 := 0; fi12 < 8; fi12++ {
                fmt.Println(fi12, fi11, v6)
                _ = fi12
            }
            fmt.Println(v1, v1)
            v6 -= (v6 >> ((2) & 63))
        } else {
            delete(v5, v1)
            v10 += float32(v3)
        }
        _ = fi11
    }
    for fi13 := 0; fi13 < 4; fi13++ {
        var v14 int16 = (v6 / (((v6 * -32768)) | 1))
        _ = fi13
        _ = v14
    }
    switch "hello" {
    case "":
        switch "hello" {
        case "hello":
            v3 = ((((-2.25 + (((float64(v10) - float64(3.14159)) + v3) + (-2.25 + v3))) + v3) + 1e10) / 3.14159)
        case "a":
            switch "foo bar" {
            case "a":
                for fi15 := 0; fi15 < 2; fi15++ {
                    switch "\n\t" {
                    case "foo bar":
                        var v16 float64 = v3
                        _ = v16
                    default:
                        var v17 float64 = (0.0 / (((float64(v10) * v3) * 1e10) * (v3 / v3)))
                        _ = v17
                    }
                    for fi18 := 0; fi18 < 9; fi18++ {
                        if (1.5 < v3) {
                            var v19 int8 = -128
                            var v20 int = 379
                            _ = v19
                            _ = v20
                        } else {
                            if true {
                                v10 += v10
                                _ = len(v5)
                                switch fi18 {
                                case -9223372036854775808:
                                    var v21 float64 = (v3 + v3)
                                    _ = v21
                                case -1:
                                    if (-128 <= 127) {
                                        v22 := GenS1{F0: 124}
                                        if false {
                                            v1 = v1
                                            switch -242 {
                                            case 4294967295:
                                                v6 -= (v6 & (v6 << ((4294967295) & 63)))
                                            }
                                        }
                                        _ = v22
                                    } else {
                                        var v23 int = fi18
                                        var v24 float32 = (float32(0.0) - v10)
                                        _ = v23
                                        _ = v24
                                    }
                                default:
                                    v25 := GenS2{F0: "\n\t", F1: ("a" + "hello"), F2: 175}
                                    _ = v25
                                }
                            } else {
                                v26 := v5[(v1 << ((4294967295) & 63))]
                                _ = v26
                            }
                        }
                        v5[454] = false
                        var v27 uint = 18446744073709551615
                        _ = fi18
                        _ = v27
                    }
                    _ = fi15
                }
            case "\n\t":
                for fi28 := 0; fi28 < 17; fi28++ {
                    v6 += v6
                    _ = fi28
                }
            case "foo bar":
                _ = len(v5)
            default:
                v1 -= (v1 ^ v1)
            }
        default:
            for fi29 := 0; fi29 < 18; fi29++ {
                switch "foo bar" {
                case "\n\t":
                    var v30 int16 = 32767
                    _ = v30
                default:
                    delete(v5, v1)
                }
                for fi31 := 0; fi31 < 6; fi31++ {
                    if (true && (3.14159 == v10)) {
                        v32 := v5[(^(v1))]
                        if len(v4) > 0 {
                            v4[0] = fi29
                        }
                        v33 := GenS3{F0: 0, F1: (v1 | 42)}
                        _ = v32
                        _ = v33
                    }
                    var v34 uint64 = (v2 % ((18446744073709551615) | 1))
                    _ = fi31
                    _ = v34
                }
                _ = fi29
            }
        }
    case "\n\t":
        var v35 int32 = 2147483647
        _ = v35
    }
    v3 = v3
    for fi36 := 0; fi36 < 8; fi36++ {
        v37 := GenS3{F0: 42, F1: (v1 - (v1 >> ((320) & 63)))}
        var v38 string = ("\n\t" + "\n\t")
        _ = len(v4)
        _ = fi36
        _ = v37
        _ = v38
    }
    fmt.Println(v1, v1)
    var v39 int16 = v6
    fmt.Println(v6)
    switch "" {
    case "hello":
        fmt.Println(v3)
    case "foo bar":
        var v40 uint = 18446744073709551615
        _ = v40
    default:
        var v41 int32 = 595
        _ = v41
    }
    for fi42 := 0; fi42 < 7; fi42++ {
        v3 = float64(0.0)
        v10 = v10
        for fi43 := 0; fi43 < 8; fi43++ {
            if true {
                if false {
                    var v44 int32 = 1
                    v45 := v5[65535]
                    _ = v44
                    _ = v45
                } else {
                    fmt.Println(v6, fi43, v10)
                }
                v10 -= v10
            } else {
                v39 = (v6 | v6)
                v3 += float64((v10 * 1e10))
                v3 += 1.5
            }
            v46 := GenS4{F0: fi42}
            _ = fi43
            _ = v46
        }
        _ = fi42
    }
    var v47 int16 = v39
    v10 += (-2.25 - ((v10 * (0.1 * (1.5 / 1e10))) + v10))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v47:", v47)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v47
    _ = v5
    _ = v6
}
