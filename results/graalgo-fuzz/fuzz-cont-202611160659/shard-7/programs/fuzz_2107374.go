package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
type GenS2 struct { F0 int64; F1 int16; F2 bool }
type GenS3 struct { F0 int; F1 string }

func main() {
    v1 := make(map[string]uint32)
    var v2 float64 = float64((float32((0.1 * 1e10)) * (1.5 - 1.5)))
    var v3 int8 = -128
    var v4 float64 = 0.1
    var v5 float32 = 3.14159
    v6 := GenS1{F0: 42, F1: -908}
    v7 := []int{}
    v4 = float64(3.14159)
    switch -2147483648 {
    case 1:
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
    case 9223372036854775807:
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
    case 128:
        var v8 string = "a"
        _ = v8
    default:
        var v9 uint = 65535
        _ = v9
    }
    if len(v7) > 0 {
        v7[0] = 128
    }
    if true {
        switch 1 {
        case 128:
            fmt.Println(v5, v3, v3)
        case 127:
            var v10 int16 = -914
            _ = v10
        }
        var v11 int8 = v3
        _ = v11
    }
    var v12 string = "\n\t"
    if (-792 == 42) {
        v13 := GenS2{F0: 0, F1: 1, F2: ((1e10 - v5) <= 3.14159)}
        _ = v13
    } else {
        if false {
            v3 = (v3 % ((v3) | 1))
            if (850 > 65535) {
                if false {
                    for fi14 := 0; fi14 < 16; fi14++ {
                        switch (v12 + v12) {
                        case "hello":
                            fmt.Println(v3, v12)
                        case "\n\t":
                            if false {
                                for fi15 := 0; fi15 < 9; fi15++ {
                                    v4 += v2
                                    v1["foo bar"] = 0
                                    _ = fi15
                                }
                            }
                        default:
                            if len(v7) >= 2 {
                                v7 = v7[1:len(v7)]
                            }
                        }
                        _ = fi14
                    }
                    for fi16 := 0; fi16 < 15; fi16++ {
                        fmt.Println(v3, v12, v3)
                        for fi17 := 0; fi17 < 13; fi17++ {
                            fmt.Println(v3, v5)
                            _ = fi17
                        }
                        switch fi16 {
                        case -164:
                            v2 = 1e10
                        default:
                            v12 = v12
                        }
                        _ = fi16
                    }
                    v18 := GenS3{F0: 32767, F1: v12}
                    _ = v18
                } else {
                    v2 = 3.14159
                    switch "foo bar" {
                    case "":
                        var v19 int16 = 32767
                        _ = v19
                    }
                }
                for fi20 := 0; fi20 < 2; fi20++ {
                    switch (fi20 << ((1) & 63)) {
                    case -1:
                        switch 128 {
                        case 928:
                            var v21 bool = true
                            _ = v21
                        case 1:
                            fmt.Println(v3)
                        default:
                            for fi22 := 0; fi22 < 15; fi22++ {
                                v23 := v1[v12]
                                _ = fi22
                                _ = v23
                            }
                        }
                    case 479:
                        if len(v7) >= 2 {
                            v7 = v7[1:len(v7)]
                        }
                    default:
                        var v24 int8 = v3
                        _ = v24
                    }
                    fmt.Println(v3, fi20, fi20)
                    _ = fi20
                }
            }
        }
        v5 = (v5 * float32(0.0))
    }
    switch (v12 + v12) {
    case "a":
        var v25 int16 = 32767
        _ = v25
    case "hello":
        fmt.Println(v4, v5)
    default:
        v26 := GenS3{F0: 2147483647, F1: ("a" + v12)}
        _ = v26
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
