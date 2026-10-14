package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8; F2 int }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint64; F1 bool }
func genFn0() int8 {
    for fi4 := 0; fi4 < 7; fi4++ {
        var v5 uint16 = 711
        v5 = v5
        _ = fi4
        _ = v5
    }
    var v6 int64 = -9223372036854775808
    v6 = (v6 + -334)
    if false {
        v6 -= -9223372036854775808
        v6 -= 42
    }
    _ = v6
    return -128
}

func main() {
    var v1 float64 = 1.5
    v2 := make(map[string]int)
    var v3 uint64 = 4294967295
    v7 := []int{-926, -759}
    var v8 uint = 895
    var v9 float32 = 3.14159
    v10 := GenS1{F0: -1, F1: 128, F2: -209}
    switch "foo bar" {
    case "a":
        _ = len(v7)
    default:
        v8 = (v8 - (v8 + 882))
    }
    v1 -= 3.14159
    var v11 float64 = (v1 + 3.14159)
    var v12 uint32 = 2147483647
    v1 -= v11
    fmt.Println(v8)
    v13 := GenS1{F0: -2147483648, F1: 255, F2: 128}
    for fi14 := 0; fi14 < 10; fi14++ {
        for fi15 := 0; fi15 < 12; fi15++ {
            if (false || ((uint16(v12) >= 65535) || (381 != 4294967295))) {
                if false {
                    if len(v7) >= 2 {
                        v7 = v7[1:len(v7)]
                    }
                }
                v1 -= 0.0
                switch fi14 {
                case 32768:
                    if false {
                        v3 += v3
                        switch fi15 {
                        case -607:
                            if (!true) {
                                switch "foo bar" {
                                case "hello":
                                    v16 := GenS1{F0: 0, F1: uint8(v8), F2: 32767}
                                    var v17 uint8 = 128
                                    if (((v3 << ((v8) & 63)) > (v3 * (^(v3)))) || (0 <= -745)) {
                                        _ = len(v2)
                                        var v18 float32 = float32((1e10 / v1))
                                        var v19 uint8 = 242
                                        if len(v7) > 0 {
                                            v7[0] = -911
                                        }
                                        _ = v18
                                        _ = v19
                                    }
                                    _ = v16
                                    _ = v17
                                case "":
                                    v7 = append(v7, fi14)
                                }
                                var v20 bool = true
                                v21 := GenS1{F0: 26, F1: 2, F2: (fi14 * fi15)}
                                _ = v20
                                _ = v21
                            }
                        }
                        fmt.Println(v12)
                    } else {
                        var v22 int32 = -1
                        _ = v22
                    }
                case -2147483648:
                    v23 := GenS2{F0: 0}
                    _ = v23
                case 952:
                    switch -706 {
                    case 32768:
                        var v24 bool = (((1.5 / 0.1) * 1e10) >= (((v11 - v1) + -2.25) - 0.1))
                        _ = v24
                    }
                }
                delete(v2, ("foo bar" + (("foo bar" + "\n\t") + "foo bar")))
            }
            var v25 float64 = (1e10 + 1.5)
            var v26 int64 = -1
            _ = fi15
            _ = v25
            _ = v26
        }
        fmt.Println(v1, fi14)
        switch "\n\t" {
        case "\n\t":
            var v27 int32 = -1
            _ = v27
        }
        _ = fi14
    }
    v2[""] = -663
    if (true && (!false)) {
        var v28 uint16 = 2
        for fi29 := 0; fi29 < 15; fi29++ {
            v3 = 136
            for fi30 := 0; fi30 < 7; fi30++ {
                var v31 string = "hello"
                v32 := GenS3{F0: uint64(v28), F1: true}
                _ = fi30
                _ = v31
                _ = v32
            }
            _ = fi29
        }
        _ = v28
    }
    v33 := GenS2{F0: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v33
    _ = v7
    _ = v8
    _ = v9
}
