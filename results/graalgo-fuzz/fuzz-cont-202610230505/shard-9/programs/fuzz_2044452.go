package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 int32 }
type GenS2 struct { F0 uint64; F1 uint64; F2 int16 }
type GenS3 struct { F0 uint64; F1 uint64 }
type GenS4 struct { F0 string; F1 uint8; F2 int16 }
func genFn0(p0 bool, p1 int8, p2 uint16) int {
    var v5 int16 = int16(p2)
    if (!p0) {
        var v6 uint16 = 170
        _ = v6
    } else {
        v7 := GenS1{F0: 65535, F1: "\n\t", F2: 608}
        _ = v7
    }
    v8 := GenS2{F0: 184, F1: 0, F2: (v5 ^ v5)}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v8
    return 32768
}

func main() {
    var v1 uint64 = 756
    v2 := make(map[uint32]int)
    v3 := make(map[string]uint32)
    var v4 uint16 = 252
    fmt.Println(v1)
    v9 := []int{}
    var v10 int64 = 231
    v10 = (-(v10))
    switch -1 {
    case 0:
        v10 += 1
    case 2147483647:
        var v11 int32 = 890
        _ = v11
    default:
        for fi12 := 0; fi12 < 18; fi12++ {
            for fi13 := 0; fi13 < 6; fi13++ {
                if ((true || (true || false)) && false) {
                    if (!((genFn0(((!((!(-2.25 < ((-2.25 - 1.5) * float32(1.5)))) || (-1 >= -1))) || true), -128, v4) <= (fi12 << ((42) & 63))) || true)) {
                        v4 = (^(v4))
                        if (127 <= 163) {
                            v14 := GenS3{F0: 734, F1: v1}
                            var v15 int16 = -568
                            var v16 uint32 = 65535
                            _ = v14
                            _ = v15
                            _ = v16
                        }
                        _ = len(v3)
                        var v17 bool = (!true)
                        v1 = (^(v1))
                        _ = v17
                    }
                } else {
                    if (((-2.25 / 1.0) / (float32((1e10 - float64(((float32((3.14159 - float64(1e10))) - 3.14159) + 3.14159)))) * 1e10)) == (1e10 / 1e10)) {
                        if ((true && false) || true) {
                            v2[628] = (fi12 % ((int(v1)) | 1))
                        }
                        v10 = v10
                    }
                    var v18 int8 = 100
                    if true {
                        switch "" {
                        case "":
                            var v19 uint32 = 1
                            _ = v19
                        case "\n\t":
                            if len(v9) > 0 {
                                v9[0] = (fi12 - fi13)
                            }
                        case "foo bar":
                            var v20 int8 = -123
                            _ = v20
                        }
                    } else {
                        v21 := GenS4{F0: "hello", F1: 1, F2: -32768}
                        if true {
                            v18 = (v18 & (v18 << ((255) & 63)))
                        }
                        fmt.Println(v4)
                        _ = v21
                    }
                    _ = v18
                }
                fmt.Println(fi12)
                _ = fi13
            }
            _ = fi12
        }
    }
    v22 := GenS4{F0: ("foo bar" + ""), F1: 127, F2: -1}
    if len(v9) > 0 {
        v9[0] = 562
    }
    v10 = v10
    v23 := GenS4{F0: "foo bar", F1: 204, F2: 0}
    v24 := GenS1{F0: 42, F1: "\n\t", F2: 680}
    var v25 bool = true
    fmt.Println(v4, v10, v4)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v9
}
