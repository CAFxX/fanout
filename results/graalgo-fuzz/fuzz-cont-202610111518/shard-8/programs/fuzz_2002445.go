package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 uint64 }
type GenS2 struct { F0 string; F1 uint }
type GenS3 struct { F0 uint32; F1 uint8; F2 bool }
type GenS4 struct { F0 uint16; F1 int8 }
type GenS5 struct { F0 int; F1 bool; F2 uint64 }
type GenS6 struct { F0 uint; F1 bool }

func main() {
    var v1 int64 = -953
    var v2 bool = false
    v3 := make(map[string]int)
    v4 := []int{42, -766, 559}
    if v2 {
        v2 = (v2 && v2)
        v1 += (v1 - (v1 / (((v1 << ((42) & 63))) | 1)))
        var v5 uint16 = 42
        _ = v5
    }
    var v6 int64 = (-(v1))
    var v7 uint64 = 187
    var v8 float64 = 0.0
    var v9 uint32 = 2147483648
    var v10 uint = 0
    for fi11 := 0; fi11 < 1; fi11++ {
        fmt.Println(v6)
        switch "hello" {
        case "\n\t":
            if (100 < -81) {
                _ = len(v3)
                v12 := GenS1{F0: "", F1: 127, F2: v7}
                var v13 uint32 = (v9 + v9)
                _ = v12
                _ = v13
            }
        case "":
            switch fi11 {
            case 2147483647:
                if v2 {
                    var v14 float32 = 1.5
                    _ = v14
                } else {
                    switch "foo bar" {
                    case "hello":
                        var v15 bool = v2
                        _ = v15
                    }
                    v10 = v10
                }
            }
        default:
            if (4294967295 > 1) {
                v16 := GenS2{F0: "a", F1: (v10 ^ v10)}
                _ = v16
            }
        }
        if v2 {
            for fi17 := 0; fi17 < 19; fi17++ {
                var v18 int64 = (v6 >> (((v10 * 911)) & 63))
                _ = fi17
                _ = v18
            }
            v4 = append(v4, fi11)
        }
        _ = fi11
    }
    v7 -= (v7 >> ((v10) & 63))
    v9 = (v9 - 2147483648)
    switch 128 {
    case 959:
        v19 := GenS3{F0: (v9 % ((42) | 1)), F1: 42, F2: v2}
        _ = v19
    case -900:
        v20 := GenS4{F0: 365, F1: int8(v9)}
        _ = v20
    case -931:
        switch "" {
        case "hello":
            for fi21 := 0; fi21 < 1; fi21++ {
                var v22 float32 = (0.0 / 1.0)
                var v23 float64 = v8
                _ = fi21
                _ = v22
                _ = v23
            }
        case "foo bar":
            fmt.Println(v1)
        case "\n\t":
            if len(v4) > 0 {
                v4[0] = -2147483648
            }
        }
    default:
        if (!(!v2)) {
            switch -546 {
            case -9223372036854775808:
                for fi24 := 0; fi24 < 2; fi24++ {
                    v8 += v8
                    switch 128 {
                    case 9223372036854775807:
                        v25 := GenS1{F0: "\n\t", F1: 42, F2: v7}
                        _ = v25
                    case -831:
                        if v2 {
                            v2 = (2 < 42)
                            if (!(true || v2)) {
                                var v26 string = "foo bar"
                                var v27 uint = v10
                                _ = v26
                                _ = v27
                            } else {
                                if (461 == 42) {
                                    fmt.Println(fi24, v9, v8)
                                    switch fi24 {
                                    case 494:
                                        var v28 uint64 = 414
                                        _ = v28
                                    case 32767:
                                        v29 := GenS2{F0: "", F1: v10}
                                        _ = v29
                                    }
                                } else {
                                    var v30 int32 = -2147483648
                                    fmt.Println(v9)
                                    if ((fi24 | fi24) >= (fi24 | (fi24 % ((fi24) | 1)))) {
                                        if v2 {
                                            fmt.Println(v30, v1)
                                            if (42 < 0) {
                                                for fi31 := 0; fi31 < 8; fi31++ {
                                                    fmt.Println(v1, v7)
                                                    _ = fi31
                                                }
                                                var v32 int16 = 804
                                                _ = v32
                                            }
                                            fmt.Println(v30)
                                        } else {
                                            for fi33 := 0; fi33 < 9; fi33++ {
                                                if v2 {
                                                    _ = len(v4)
                                                    v34 := GenS4{F0: 2, F1: 100}
                                                    _ = v34
                                                } else {
                                                    if (!(v30 > (-(v30)))) {
                                                        v8 += v8
                                                        v35 := GenS3{F0: 4294967295, F1: 1, F2: false}
                                                        fmt.Println(v1)
                                                        _ = v35
                                                    }
                                                }
                                                _ = fi33
                                            }
                                        }
                                        v36 := GenS5{F0: (fi24 * fi24), F1: false, F2: v7}
                                        v10 = 1
                                        _ = v36
                                    }
                                    _ = v30
                                }
                                v37 := GenS6{F0: (v10 << ((v10) & 63)), F1: v2}
                                _ = v37
                            }
                        } else {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        }
                    }
                    _ = fi24
                }
            case 788:
                fmt.Println(v10, v1, v1)
            default:
                v6 = (v1 / (((v1 ^ (v6 >> (((v10 >> (((v10 % ((v10) | 1))) & 63))) & 63)))) | 1))
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
