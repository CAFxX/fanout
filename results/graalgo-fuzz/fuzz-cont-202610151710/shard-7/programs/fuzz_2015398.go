package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
type GenS2 struct { F0 int; F1 uint }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint16; F1 uint8 }
type GenS6 struct { F0 uint64; F1 uint; F2 int16 }
type GenS7 struct { F0 int16; F1 uint; F2 int32 }
type GenS8 struct { F0 string; F1 uint16; F2 int8 }
type GenS9 struct { F0 bool }

func main() {
    v1 := make(map[int]bool)
    v2 := []int{127, -2147483648, -2147483648, -1}
    v3 := make(map[int]int)
    var v4 uint32 = 1
    var v5 uint32 = (v4 - 843)
    v5 += (v4 | 783)
    v6 := GenS1{F0: -1, F1: -1}
    v5 = v4
    fmt.Println(v4)
    v2 = append(v2, -1)
    v4 = 4294967295
    delete(v3, -311)
    for fi7 := 0; fi7 < 15; fi7++ {
        var v8 int16 = 42
        _ = fi7
        _ = v8
    }
    if true {
        v4 = (v5 | v5)
        switch -1 {
        case -9223372036854775808:
            _ = len(v3)
        }
        for fi9 := 0; fi9 < 9; fi9++ {
            if false {
                if false {
                    fmt.Println(v4, v5)
                    v10 := GenS1{F0: 0, F1: -848}
                    if len(v2) > 0 {
                        v2[0] = fi9
                    }
                    _ = v10
                } else {
                    if true {
                        v4 = v5
                        for fi11 := 0; fi11 < 8; fi11++ {
                            var v12 float32 = (0.1 * -2.25)
                            v13 := GenS2{F0: fi11, F1: 735}
                            fmt.Println(v5)
                            fmt.Println(fi9)
                            _ = fi11
                            _ = v12
                            _ = v13
                        }
                        v5 = (v5 << ((42) & 63))
                    } else {
                        if len(v2) > 0 {
                            v2[0] = (fi9 >> ((127) & 63))
                        }
                        switch "hello" {
                        case "foo bar":
                            v14 := GenS2{F0: fi9, F1: 42}
                            _ = v14
                        case "":
                            _ = len(v1)
                        default:
                            v15 := GenS3{F0: 1}
                            _ = v15
                        }
                        v5 = v4
                    }
                }
                _ = len(v2)
                delete(v3, fi9)
            } else {
                v5 -= v4
                switch (fi9 % (((fi9 % (((fi9 / ((458) | 1))) | 1))) | 1)) {
                case 128:
                    for fi16 := 0; fi16 < 16; fi16++ {
                        if (915 <= 1) {
                            if ((((3.14159 * (1.5 * (-2.25 / (3.14159 - 1e10)))) - 0.1) - 1.5) != 0.1) {
                                var v17 float32 = 3.14159
                                switch 32767 {
                                case 831:
                                    v4 = 4294967295
                                case 32767:
                                    v4 = v5
                                case -9223372036854775808:
                                    var v18 uint16 = 1
                                    _ = v18
                                }
                                _ = v17
                            } else {
                                v2 = append(v2, fi16)
                                v5 = (v4 >> ((2) & 63))
                                var v19 int32 = -1
                                _ = v19
                            }
                            v5 = (^(v4))
                            var v20 uint = 2
                            _ = v20
                        }
                        if (true || false) {
                            var v21 int32 = -2147483648
                            _ = v21
                        }
                        _ = fi16
                    }
                case 881:
                    switch (("hello" + "a") + "") {
                    case "foo bar":
                        if (608 >= -1) {
                            _ = len(v2)
                            switch fi9 {
                            case 128:
                                v22 := GenS4{F0: 45}
                                _ = v22
                            default:
                                if ((v4 % ((v4) | 1)) == 0) {
                                    var v23 uint = 42
                                    var v24 int8 = -84
                                    for fi25 := 0; fi25 < 14; fi25++ {
                                        if (0.1 >= (0.1 - ((0.1 * (float32(((-2.25 - 3.14159) - (0.0 - -2.25))) / 0.1)) + 1.5))) {
                                            v5 -= (v5 ^ (v5 - v4))
                                            v26 := GenS1{F0: v24, F1: 565}
                                            _ = v26
                                        } else {
                                            switch "\n\t" {
                                            case "foo bar":
                                                v3[fi9] = (-(fi9))
                                            case "\n\t":
                                                if true {
                                                    if (!true) {
                                                        v23 += v23
                                                        v4 = (v5 << ((4294967295) & 63))
                                                        var v27 int16 = 1000
                                                        _ = v27
                                                    }
                                                    switch "" {
                                                    case "":
                                                        if false {
                                                            var v28 int8 = (v24 / (((v24 - v24)) | 1))
                                                            fmt.Println(v28)
                                                            v29 := GenS5{F0: 300, F1: 1}
                                                            _ = v28
                                                            _ = v29
                                                        }
                                                    }
                                                    v30 := GenS1{F0: -1, F1: -1}
                                                    _ = v30
                                                }
                                            }
                                            if (false && true) {
                                                _ = len(v2)
                                                v23 = 1
                                                v3[fi9] = (fi25 / ((fi9) | 1))
                                            } else {
                                                var v31 int = (fi9 ^ (fi9 ^ (fi25 % (((fi9 * (fi25 - fi9))) | 1))))
                                                v2 = append(v2, -9223372036854775808)
                                                _ = v31
                                            }
                                            v32 := GenS6{F0: 1, F1: 0, F2: -32768}
                                            var v33 float32 = 0.0
                                            var v34 uint8 = 127
                                            _ = v32
                                            _ = v33
                                            _ = v34
                                        }
                                        if len(v2) > 0 {
                                            v2[0] = 4294967295
                                        }
                                        _ = fi25
                                    }
                                    _ = v23
                                    _ = v24
                                }
                            }
                        } else {
                            var v35 int8 = -1
                            fmt.Println(fi9, v4, v4)
                            _ = v35
                        }
                    case "hello":
                        if (!(true && false)) {
                            fmt.Println(fi9)
                            v5 = (v5 & (v4 / (((v4 >> ((4294967295) & 63))) | 1)))
                            v36 := GenS2{F0: fi9, F1: 55}
                            _ = v36
                        } else {
                            v37 := GenS5{F0: uint16(v5), F1: 167}
                            _ = v37
                        }
                    case "\n\t":
                        var v38 int16 = -32768
                        _ = v38
                    default:
                        fmt.Println(v5, v5, v5)
                    }
                }
            }
            v39 := GenS3{F0: 11}
            _ = fi9
            _ = v39
        }
    }
    switch "\n\t" {
    case "a":
        v40 := GenS7{F0: -1, F1: 42, F2: 1}
        _ = v40
    case "":
        if (!true) {
            v41 := GenS8{F0: "\n\t", F1: 2, F2: -128}
            switch 127 {
            case 4294967295:
                fmt.Println(v5)
            }
            _ = v41
        }
    case "foo bar":
        fmt.Println(v5)
    default:
        var v42 uint32 = (^(v4))
        _ = v42
    }
    var v43 uint = 544
    var v44 float64 = 0.1
    var v45 int32 = -139
    var v46 float32 = 1e10
    v43 = (v43 % ((113) | 1))
    fmt.Println(v5, v46, v46)
    if false {
        _ = len(v2)
    } else {
        v47 := GenS9{F0: false}
        var v48 int32 = v45
        _ = v47
        _ = v48
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v5
    _ = v6
}
