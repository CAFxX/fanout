package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64; F1 int64; F2 string }
type GenS3 struct { F0 int32; F1 uint; F2 string }
type GenS4 struct { F0 uint32; F1 uint }

func main() {
    v1 := make(map[string]bool)
    v2 := []int{0}
    var v3 int8 = -91
    v4 := make(map[uint32]int)
    var v5 float32 = 1e10
    v6 := make(map[int]int)
    fmt.Println(v5, v5)
    var v7 float32 = float32((3.14159 * float64(0.0)))
    v7 = (v7 - ((((0.0 * 3.14159) * v7) / v5) + 1.5))
    for fi8 := 0; fi8 < 9; fi8++ {
        if (true || (42 <= 128)) {
            fmt.Println(fi8, v3, fi8)
        } else {
            v4[65535] = 601
        }
        _ = fi8
    }
    if (-122 >= 32767) {
        v5 += (float32((-2.25 - (float64(v5) - 3.14159))) / v5)
    }
    if false {
        _ = len(v2)
    }
    for fi9 := 0; fi9 < 4; fi9++ {
        var v10 int64 = -1
        fmt.Println(fi9, v3, v3)
        _ = fi9
        _ = v10
    }
    var v11 uint8 = 0
    for fi12 := 0; fi12 < 9; fi12++ {
        var v13 int16 = 249
        if true {
            var v14 int32 = 385
            _ = len(v2)
            _ = v14
        } else {
            v15 := v6[(fi12 + fi12)]
            if (-1 != -2147483648) {
                v16 := GenS1{F0: 1}
                _ = v16
            }
            v11 = v11
            _ = v15
        }
        for fi17 := 0; fi17 < 7; fi17++ {
            switch -2147483648 {
            case 390:
                if true {
                    switch "hello" {
                    case "foo bar":
                        if (fi17 != (fi12 >> ((2) & 63))) {
                            if true {
                                fmt.Println(v7, v11, v11)
                                v5 -= 0.1
                                switch 9223372036854775807 {
                                case 32767:
                                    fmt.Println(v5, v11, fi12)
                                }
                                v18 := GenS1{F0: 1}
                                _ = v18
                            }
                            switch fi12 {
                            case 564:
                                fmt.Println(fi12)
                            default:
                                if true {
                                    if (128 != fi17) {
                                        v19 := GenS2{F0: 138, F1: 9223372036854775807, F2: ""}
                                        v20 := GenS3{F0: -1, F1: 1, F2: "foo bar"}
                                        var v21 float32 = ((v5 + 3.14159) / -2.25)
                                        _ = v19
                                        _ = v20
                                        _ = v21
                                    } else {
                                        v5 = v7
                                    }
                                } else {
                                    _ = len(v2)
                                    if false {
                                        fmt.Println(v13)
                                        if false {
                                            fmt.Println(fi17)
                                            var v22 uint64 = 620
                                            var v23 int16 = v13
                                            _ = v22
                                            _ = v23
                                        } else {
                                            var v24 bool = true
                                            _ = v24
                                        }
                                        var v25 int64 = -77
                                        _ = v25
                                    }
                                }
                            }
                            v3 = (v3 ^ 0)
                        }
                    case "hello":
                        v7 -= 1.5
                    default:
                        switch (fi12 << ((42) & 63)) {
                        case 1:
                            fmt.Println(v5, v3, v13)
                        case 4294967295:
                            var v26 uint64 = 232
                            _ = v26
                        case 88:
                            var v27 uint16 = 896
                            _ = v27
                        }
                    }
                }
            case -2147483648:
                v28 := v4[330]
                _ = v28
            default:
                if true {
                    v7 = 3.14159
                    v29 := GenS1{F0: 2}
                    v7 = 0.1
                    _ = v29
                }
            }
            var v30 uint32 = 0
            _ = fi17
            _ = v30
        }
        _ = fi12
        _ = v13
    }
    v3 -= (-(v3))
    if (true && ((v11 ^ (v11 | v11)) <= v11)) {
        v31 := GenS2{F0: 711, F1: 442, F2: "foo bar"}
        var v32 uint64 = 2
        v33 := GenS4{F0: 465, F1: 0}
        _ = v31
        _ = v32
        _ = v33
    }
    var v34 uint64 = 923
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
