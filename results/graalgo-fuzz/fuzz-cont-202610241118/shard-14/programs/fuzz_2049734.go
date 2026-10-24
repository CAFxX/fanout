package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 uint16; F1 string }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint8; F1 uint }
type GenS5 struct { F0 uint8 }
func genFn0() string {
    switch "foo bar" {
    case "a":
        if (2 >= 1) {
            v6 := GenS1{F0: "a", F1: 32767}
            _ = v6
        }
    case "foo bar":
        for fi7 := 0; fi7 < 18; fi7++ {
            v8 := []int{(fi7 | (fi7 - 42)), 127, (fi7 ^ fi7), 127}
            for fi9 := 0; fi9 < 2; fi9++ {
                v8 = append(v8, (fi7 | (fi7 & (fi9 | -9223372036854775808))))
                if (-405 < -726) {
                    switch "hello" {
                    case "a":
                        var v10 bool = (0.0 <= float64((float32((-2.25 / 1.5)) / (float32((0.0 * 1.5)) - 0.1))))
                        _ = v10
                    }
                    v11 := GenS2{F0: 42, F1: "a"}
                    _ = v11
                } else {
                    var v12 int16 = 579
                    if true {
                        v13 := GenS2{F0: 42, F1: "foo bar"}
                        _ = v13
                    } else {
                        if true {
                            v12 += v12
                            _ = len(v8)
                        } else {
                            if (!true) {
                                var v14 float64 = 3.14159
                                v15 := GenS3{F0: true}
                                switch -541 {
                                case 4294967295:
                                    v14 = v14
                                case -1:
                                    switch 487 {
                                    case 32768:
                                        switch 4294967295 {
                                        case 127:
                                            v14 += v14
                                        default:
                                            v16 := GenS3{F0: (!true)}
                                            _ = v16
                                        }
                                    case -9223372036854775808:
                                        fmt.Println(fi9, v12, v14)
                                    case 32767:
                                        switch fi7 {
                                        case 4294967295:
                                            fmt.Println(v12, v14)
                                        default:
                                            v12 = v12
                                        }
                                    }
                                case 677:
                                    fmt.Println(v12)
                                default:
                                    v8 = append(v8, -9223372036854775808)
                                }
                                _ = v14
                                _ = v15
                            } else {
                                if (-1 == int32(v12)) {
                                    if len(v8) >= 2 {
                                        v8 = v8[1:len(v8)]
                                    }
                                    if true {
                                        if len(v8) > 0 {
                                            v8[0] = (fi7 % ((1) | 1))
                                        }
                                        if len(v8) > 0 {
                                            v8[0] = fi7
                                        }
                                    } else {
                                        var v17 uint64 = 4294967295
                                        _ = v17
                                    }
                                    if false {
                                        v18 := GenS3{F0: false}
                                        var v19 int16 = v12
                                        _ = v18
                                        _ = v19
                                    } else {
                                        _ = len(v8)
                                        v20 := GenS2{F0: 99, F1: "a"}
                                        v12 += 0
                                        _ = v20
                                    }
                                }
                            }
                            var v21 bool = false
                            _ = v21
                        }
                        switch 233 {
                        case 19:
                            var v22 uint64 = 0
                            _ = v22
                        default:
                            var v23 int64 = 42
                            _ = v23
                        }
                    }
                    if len(v8) > 0 {
                        v8[0] = -18
                    }
                    var v24 int8 = 127
                    _ = v12
                    _ = v24
                }
                fmt.Println(fi7)
                _ = fi9
            }
            _ = fi7
            _ = v8
        }
    case "":
        var v25 float64 = 1.5
        _ = v25
    }
    var v26 uint8 = 0
    _ = v26
    return ""
}

func main() {
    v1 := make(map[string]string)
    var v2 float64 = 3.14159
    var v3 uint32 = 686
    v4 := make(map[int]string)
    var v5 int = 1
    fmt.Println(v2)
    var v27 uint32 = (^(v3))
    var v28 float32 = 1.5
    v29 := GenS4{F0: 0, F1: 42}
    fmt.Println(v28, v2, v5)
    for fi30 := 0; fi30 < 4; fi30++ {
        v31 := GenS3{F0: (!(!false))}
        v32 := GenS4{F0: 94, F1: 42}
        _ = fi30
        _ = v31
        _ = v32
    }
    v5 -= (v5 ^ 1)
    v27 = (v3 + (v27 | 2147483648))
    fmt.Println(v28)
    switch (^(v5)) {
    case 195:
        fmt.Println(v5, v28, v3)
    case -9223372036854775808:
        var v33 int32 = 0
        _ = v33
    default:
        switch "a" {
        case "a":
            for fi34 := 0; fi34 < 9; fi34++ {
                v27 += 2
                v35 := GenS5{F0: 255}
                v27 += v3
                _ = fi34
                _ = v35
            }
        case "\n\t":
            v2 -= 1e10
        case "hello":
            v27 += v27
        default:
            fmt.Println(v27, v5)
        }
    }
    var v36 uint32 = (v3 ^ (^(v27)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v36
    _ = v4
    _ = v5
}
