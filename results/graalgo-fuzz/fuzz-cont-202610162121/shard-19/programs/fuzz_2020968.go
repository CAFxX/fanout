package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 int64 }
type GenS2 struct { F0 int32; F1 int32 }
type GenS3 struct { F0 uint }
func genFn0(p0 string, p1 uint, p2 int32) uint16 {
    var v6 float64 = (0.0 + float64((1e10 * 0.1)))
    v6 += (v6 / v6)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 749
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[string]bool)
    var v3 float32 = 0.1
    v4 := []int{-2147483648}
    var v5 uint = 463
    if (!false) {
        for fi7 := 0; fi7 < 18; fi7++ {
            var v8 int8 = 42
            var v9 int32 = -2147483648
            fmt.Println(v5, v8, v8)
            _ = fi7
            _ = v8
            _ = v9
        }
        var v10 int16 = -1
        _ = v10
    } else {
        v11 := GenS1{F0: -22, F1: v5, F2: -1}
        v12 := GenS1{F0: 2147483647, F1: (^(v5)), F2: 0}
        for fi13 := 0; fi13 < 4; fi13++ {
            switch "foo bar" {
            case "\n\t":
                var v14 uint32 = 2147483648
                _ = v14
            case "hello":
                fmt.Println(v5, fi13, v3)
            case "a":
                var v15 uint32 = 841
                _ = v15
            default:
                switch "" {
                case "\n\t":
                    switch "\n\t" {
                    case "a":
                        var v16 string = "\n\t"
                        _ = v16
                    case "":
                        if (2 > 2) {
                            switch (fi13 & fi13) {
                            case 32768:
                                var v17 int16 = 42
                                _ = v17
                            default:
                                var v18 float64 = 1e10
                                _ = v18
                            }
                            if (32767 == 255) {
                                var v19 int16 = 42
                                switch ("\n\t" + "") {
                                case "\n\t":
                                    v20 := GenS2{F0: -1, F1: -947}
                                    _ = v20
                                case "":
                                    var v21 uint8 = 255
                                    _ = v21
                                }
                                _ = v19
                            }
                        }
                    default:
                        var v22 float32 = (3.14159 * (-2.25 / ((v3 - ((1.5 - (float32(1e10) / (v3 / (v3 * (v3 * 1e10))))) * v3)) / v3)))
                        _ = v22
                    }
                default:
                    v3 += (1.5 - 0.1)
                }
            }
            var v23 uint8 = 128
            _ = fi13
            _ = v23
        }
        _ = v11
        _ = v12
    }
    for fi24 := 0; fi24 < 5; fi24++ {
        for fi25 := 0; fi25 < 11; fi25++ {
            v26 := GenS1{F0: (fi24 & fi25), F1: (v5 % ((v5) | 1)), F2: 1}
            _ = fi25
            _ = v26
        }
        var v27 uint16 = 86
        _ = fi24
        _ = v27
    }
    v3 += float32(3.14159)
    v28 := GenS3{F0: v5}
    var v29 int32 = 428
    v5 += v5
    _ = len(v4)
    if false {
        var v30 int16 = 965
        v31 := GenS1{F0: -946, F1: 0, F2: int64(v5)}
        var v32 float64 = (0.0 / -2.25)
        _ = v30
        _ = v31
        _ = v32
    }
    for fi33 := 0; fi33 < 8; fi33++ {
        var v34 uint64 = 983
        v3 = (((float32((1e10 + 0.1)) + v3) - v3) - float32((0.1 / (float64(v3) * (((((1.5 + float64(v3)) - 1.5) + float64(0.0)) * (3.14159 + -2.25)) * 3.14159)))))
        v5 = v5
        _ = fi33
        _ = v34
    }
    var v35 uint64 = 2
    for fi36 := 0; fi36 < 14; fi36++ {
        v37 := GenS3{F0: uint(v35)}
        _ = fi36
        _ = v37
    }
    var v38 uint = (v5 & (v5 / ((65535) | 1)))
    var v39 int32 = (v29 | (^(v29)))
    v3 -= (v3 + (v3 + (float32(1e10) + v3)))
    v3 = 3.14159
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v35
    _ = v38
    _ = v39
    _ = v4
    _ = v5
}
