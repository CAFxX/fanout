package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 int32 }
type GenS2 struct { F0 int16; F1 bool; F2 uint8 }
func genFn0(p0 int8, p1 uint, p2 int) int64 {
    v6 := []int{-9223372036854775808, 32767, (p2 * (p2 / (((p2 ^ p2)) | 1)))}
    fmt.Println(p1, p1, p2)
    for fi7 := 0; fi7 < 10; fi7++ {
        var v8 string = "\n\t"
        if (p2 < -2147483648) {
            switch (fi7 >> (((^(p1))) & 63)) {
            case -9223372036854775808:
                switch "hello" {
                case "":
                    var v9 int32 = 0
                    _ = v9
                case "\n\t":
                    if (!true) {
                        fmt.Println(p0)
                    }
                default:
                    var v10 float64 = (-2.25 - 0.1)
                    _ = v10
                }
            case 42:
                v11 := GenS1{F0: 0, F1: p1, F2: -1}
                _ = v11
            default:
                if (!false) {
                    var v12 bool = false
                    _ = v12
                } else {
                    for fi13 := 0; fi13 < 14; fi13++ {
                        p0 -= int8(fi7)
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        _ = fi13
                    }
                    if len(v6) > 0 {
                        v6[0] = fi7
                    }
                    var v14 uint8 = 0
                    _ = v14
                }
            }
            var v15 uint64 = 751
            fmt.Println(fi7, fi7, v15)
            _ = v15
        }
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 0
}

func main() {
    v1 := []uint32{}
    var v2 bool = true
    var v3 uint64 = 269
    var v4 int32 = 628
    var v5 string = ""
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v16 := GenS1{F0: v3, F1: 491, F2: (v4 ^ v4)}
    for fi17 := 0; fi17 < 19; fi17++ {
        var v18 uint64 = (v3 + 18446744073709551615)
        _ = fi17
        _ = v18
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if (v2 && false) {
        var v19 int16 = -831
        _ = v19
    }
    fmt.Println(v2, v5)
    v5 = v5
    switch "" {
    case "\n\t":
        fmt.Println(v2)
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi20 := 0; fi20 < 2; fi20++ {
        if (v2 && false) {
            v21 := GenS1{F0: 503, F1: 42, F2: (v4 ^ v4)}
            for fi22 := 0; fi22 < 9; fi22++ {
                if true {
                    v3 += 0
                    var v23 float32 = 0.0
                    v5 = v5
                    _ = v23
                } else {
                    _ = len(v1)
                    v1 = append(v1, 2147483648)
                    var v24 float32 = -2.25
                    _ = v24
                }
                if v2 {
                    v3 = 385
                } else {
                    _ = len(v1)
                    fmt.Println(v2, v2, v5)
                    v25 := GenS2{F0: 485, F1: v2, F2: 2}
                    var v26 float32 = -2.25
                    _ = v25
                    _ = v26
                }
                _ = fi22
            }
            _ = v21
        } else {
            v3 -= v3
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        fmt.Println(v3)
        _ = fi20
    }
    for fi27 := 0; fi27 < 1; fi27++ {
        if len(v1) > 0 {
            v1[0] = 42
        }
        _ = fi27
    }
    var v28 float32 = float32(((0.1 + 0.1) * float64(-2.25)))
    var v29 uint32 = 847
    var v30 uint = 255
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
