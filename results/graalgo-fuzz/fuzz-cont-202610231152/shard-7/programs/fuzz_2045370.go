package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64; F1 uint8 }
type GenS3 struct { F0 uint32; F1 uint64 }
func genFn0(p0 uint64, p1 uint32) uint8 {
    p0 = (p0 + p0)
    var v5 int64 = 0
    if (!(!(0 >= 1))) {
        fmt.Println(p1)
    }
    _ = p0
    _ = p1
    _ = v5
    return 241
}
func genFn1(p0 string, p1 uint32, p2 string) bool {
    p0 = "hello"
    if false {
        var v6 int64 = 42
        var v7 int = -9223372036854775808
        fmt.Println(v6, p2, p0)
        _ = v6
        _ = v7
    } else {
        var v8 int32 = -2147483648
        switch p0 {
        case "hello":
            for fi9 := 0; fi9 < 20; fi9++ {
                for fi10 := 0; fi10 < 14; fi10++ {
                    v11 := []int{(fi9 / ((fi9) | 1))}
                    switch fi9 {
                    case 4294967295:
                        if len(v11) > 0 {
                            v11[0] = -9223372036854775808
                        }
                    case 42:
                        fmt.Println(v8)
                    }
                    _ = fi10
                    _ = v11
                }
                _ = fi9
            }
        case "a":
            var v12 float64 = (-2.25 * ((0.0 * 0.1) * (float64(-2.25) * 0.0)))
            _ = v12
        case "":
            var v13 float32 = float32(0.0)
            _ = v13
        default:
            fmt.Println(p0, p0, p0)
        }
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    return (false || (291 <= 1))
}

func main() {
    var v1 uint32 = 2
    v2 := make(map[int]uint32)
    var v3 int64 = 0
    var v4 int64 = (v3 | v3)
    var v14 int32 = 0
    v15 := []int{128}
    v4 += 316
    v16 := GenS1{F0: v3}
    fmt.Println(v1)
    var v17 uint = 4294967295
    fmt.Println(v3, v17, v14)
    fmt.Println(v17)
    for fi18 := 0; fi18 < 8; fi18++ {
        fmt.Println(v14, v4)
        _ = fi18
    }
    switch 42 {
    case 1:
        v17 -= (v17 - (v17 << (((^(v17))) & 63)))
    case 271:
        if (genFn0(uint64(v1), v1) < 42) {
            var v19 int = 128
            _ = v19
        } else {
            v20 := GenS1{F0: (v3 & 9223372036854775807)}
            _ = v20
        }
    }
    switch 9223372036854775807 {
    case 4294967295:
        if (!(157 <= 75)) {
            v15 = append(v15, 128)
            for fi21 := 0; fi21 < 20; fi21++ {
                if false {
                    if (-400 == fi21) {
                        switch "hello" {
                        case "hello":
                            switch (^(fi21)) {
                            case -1:
                                for fi22 := 0; fi22 < 4; fi22++ {
                                    v14 = v14
                                    _ = fi22
                                }
                            case -284:
                                var v23 int32 = v14
                                _ = v23
                            default:
                                v4 -= -1
                            }
                        case "":
                            v24 := GenS1{F0: v3}
                            _ = v24
                        }
                        var v25 int = int(v17)
                        _ = v25
                    } else {
                        for fi26 := 0; fi26 < 5; fi26++ {
                            switch "\n\t" {
                            case "\n\t":
                                fmt.Println(v14)
                            }
                            var v27 uint = (v17 % ((2) | 1))
                            _ = fi26
                            _ = v27
                        }
                        var v28 float64 = 0.0
                        for fi29 := 0; fi29 < 19; fi29++ {
                            v28 = v28
                            v30 := GenS1{F0: v4}
                            _ = fi29
                            _ = v30
                        }
                        _ = v28
                    }
                }
                var v31 uint = (^(v17))
                for fi32 := 0; fi32 < 14; fi32++ {
                    if false {
                        var v33 bool = (1 != 127)
                        v34 := GenS2{F0: (v4 >> (((v31 / (((^(v17))) | 1))) & 63)), F1: 0}
                        _ = v33
                        _ = v34
                    } else {
                        var v35 int = (fi32 ^ fi21)
                        var v36 int = (v35 % (((fi21 - fi21)) | 1))
                        _ = v35
                        _ = v36
                    }
                    _ = fi32
                }
                _ = fi21
                _ = v31
            }
        } else {
            if false {
                if false {
                    _ = len(v15)
                }
            }
            for fi37 := 0; fi37 < 10; fi37++ {
                var v38 string = "a"
                _ = fi37
                _ = v38
            }
            var v39 uint = v17
            _ = v39
        }
    case -1:
        v17 = (v17 % (((v17 + (^(v17)))) | 1))
    }
    v40 := GenS3{F0: v1, F1: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v40
}
