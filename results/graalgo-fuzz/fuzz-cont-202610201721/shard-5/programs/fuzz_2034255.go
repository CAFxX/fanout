package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 int32; F1 uint32 }
type GenS3 struct { F0 uint64; F1 int16 }
type GenS4 struct { F0 int; F1 bool }
type GenS5 struct { F0 uint64; F1 uint16; F2 uint16 }
type GenS6 struct { F0 uint8; F1 int }

func main() {
    var v1 int = 1
    v2 := make(map[string]int)
    var v3 uint16 = 162
    v4 := make(map[string]int)
    var v5 float32 = 0.1
    for fi6 := 0; fi6 < 14; fi6++ {
        fmt.Println(v1)
        _ = fi6
    }
    switch (v1 << ((4294967295) & 63)) {
    case 32767:
        var v7 uint8 = 217
        _ = v7
    }
    switch ("hello" + "\n\t") {
    case "hello":
        fmt.Println(v3, v5, v1)
    case "foo bar":
        var v8 int32 = -326
        _ = v8
    case "\n\t":
        v9 := GenS1{F0: "foo bar", F1: 2147483648}
        _ = v9
    default:
        v10 := GenS2{F0: -1, F1: 1}
        _ = v10
    }
    var v11 uint8 = 205
    var v12 int64 = -9223372036854775808
    v13 := []int{(^(v1)), v1, v1}
    v5 += v5
    v3 += (v3 | v3)
    for fi14 := 0; fi14 < 18; fi14++ {
        var v15 bool = false
        if (v5 >= ((-2.25 - (1e10 / v5)) / ((float32(1e10) / -2.25) - 0.1))) {
            var v16 uint = 879
            _ = v16
        } else {
            switch ("foo bar" + "a") {
            case "foo bar":
                v1 = (-(v1))
            case "\n\t":
                if (1 != v11) {
                    if true {
                        for fi17 := 0; fi17 < 8; fi17++ {
                            if (!(-9223372036854775808 == v1)) {
                                fmt.Println(fi17, fi17, v12)
                                if len(v13) >= 2 {
                                    v13 = v13[1:len(v13)]
                                }
                                v18 := GenS1{F0: "foo bar", F1: 4294967295}
                                v19 := GenS3{F0: 0, F1: 1}
                                _ = v18
                                _ = v19
                            } else {
                                v3 = (v3 & (^(v3)))
                            }
                            v20 := GenS1{F0: "a", F1: 42}
                            fmt.Println(v12, fi17)
                            if false {
                                fmt.Println(fi14, v1)
                                v5 -= float32(float64(((-2.25 / 1.0) * (float32((float64(v5) - 3.14159)) + v5))))
                            }
                            _ = fi17
                            _ = v20
                        }
                        if v15 {
                            _ = len(v13)
                            fmt.Println(v12, v15)
                            var v21 uint = 1
                            _ = v21
                        }
                        var v22 uint64 = 4294967295
                        _ = v22
                    }
                    var v23 int32 = 1
                    _ = v23
                } else {
                    if v15 {
                        var v24 int64 = -9223372036854775808
                        v12 = v12
                        _ = v24
                    } else {
                        if (42 > 2147483647) {
                            v11 -= (v11 + 233)
                            var v25 int16 = 32
                            _ = v25
                        } else {
                            for fi26 := 0; fi26 < 5; fi26++ {
                                var v27 int8 = -128
                                var v28 uint16 = 518
                                _ = fi26
                                _ = v27
                                _ = v28
                            }
                            for fi29 := 0; fi29 < 9; fi29++ {
                                var v30 float32 = (v5 / v5)
                                _ = fi29
                                _ = v30
                            }
                            for fi31 := 0; fi31 < 12; fi31++ {
                                _ = len(v13)
                                if len(v13) > 0 {
                                    v13[0] = fi14
                                }
                                _ = fi31
                            }
                        }
                        v11 = (v11 << ((4294967295) & 63))
                        v15 = (0 < 538)
                    }
                    var v32 float64 = 0.0
                    _ = v32
                }
            case "":
                switch "hello" {
                case "hello":
                    v3 += (v3 >> ((65535) & 63))
                case "\n\t":
                    var v33 uint64 = 789
                    _ = v33
                }
            }
            if (v15 && ((42 < 18446744073709551615) || (!(((v11 % ((v11) | 1)) <= (v11 >> ((1) & 63))) || (!v15))))) {
                v34 := GenS4{F0: v1, F1: (2 <= 1)}
                switch -1 {
                case 0:
                    var v35 uint16 = (v3 / ((1) | 1))
                    _ = v35
                case -342:
                    if true {
                        switch "a" {
                        case "":
                            v36 := GenS3{F0: 42, F1: -1}
                            _ = v36
                        default:
                            v13 = append(v13, 783)
                        }
                        switch v1 {
                        case 4294967295:
                            if (v11 < (v11 & (^(v11)))) {
                                var v37 uint16 = (v3 >> ((uint(v3)) & 63))
                                _ = v37
                            } else {
                                v38 := GenS5{F0: 4294967295, F1: v3, F2: v3}
                                var v39 uint = 42
                                _ = v38
                                _ = v39
                            }
                        }
                        _ = len(v13)
                    } else {
                        v40 := GenS3{F0: 1, F1: -32768}
                        v41 := GenS6{F0: 0, F1: v1}
                        if v15 {
                            v15 = (42 != -299)
                            v12 = (v12 ^ (v12 / (((-(v12))) | 1)))
                            if v15 {
                                var v42 uint32 = 2147483647
                                fmt.Println(v11, v1, v11)
                                _ = v42
                            }
                        }
                        _ = v40
                        _ = v41
                    }
                }
                var v43 uint8 = v11
                _ = v34
                _ = v43
            }
            _ = len(v13)
        }
        _ = fi14
        _ = v15
    }
    for fi44 := 0; fi44 < 3; fi44++ {
        v11 -= v11
        var v45 int = (v1 << ((1) & 63))
        var v46 uint8 = (v11 % (((v11 * (v11 | v11))) | 1))
        _ = fi44
        _ = v45
        _ = v46
    }
    if len(v13) >= 2 {
        v13 = v13[1:len(v13)]
    }
    v3 += (v3 - v3)
    v12 += (-(v12))
    v11 -= (v11 * (v11 ^ v11))
    v12 = 143
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
