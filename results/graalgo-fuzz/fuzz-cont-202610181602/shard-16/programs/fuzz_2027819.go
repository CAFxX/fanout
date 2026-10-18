package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int32; F1 uint; F2 int32 }
type GenS4 struct { F0 int32; F1 uint16 }

func main() {
    var v1 int16 = 32767
    var v2 uint32 = 4294967295
    v3 := make(map[int]int)
    var v4 float64 = 3.14159
    var v5 bool = (1 != -9223372036854775808)
    fmt.Println(v4, v1)
    if v5 {
        var v6 uint = 1
        if (2147483647 <= 2147483647) {
            if true {
                fmt.Println(v2)
                v1 = v1
            }
        } else {
            if (32767 > int16(v6)) {
                v2 -= 0
                var v7 int64 = 1
                _ = v7
            } else {
                var v8 uint = v6
                for fi9 := 0; fi9 < 19; fi9++ {
                    if ((v2 << (((v8 / (((^(v6))) | 1))) & 63)) < v2) {
                        if v5 {
                            var v10 int32 = 1
                            _ = len(v3)
                            var v11 uint64 = 488
                            _ = v10
                            _ = v11
                        } else {
                            v12 := []int{fi9, 0}
                            for fi13 := 0; fi13 < 3; fi13++ {
                                var v14 int64 = 9223372036854775807
                                if v5 {
                                    v14 = (v14 >> (((v8 >> ((v6) & 63))) & 63))
                                    v15 := GenS1{F0: (v1 * (v1 + (v1 | v1))), F1: -675}
                                    v8 = v6
                                    _ = v15
                                } else {
                                    v8 = v8
                                }
                                switch "" {
                                case "foo bar":
                                    v6 = 0
                                case "\n\t":
                                    if false {
                                        var v16 string = ("" + "\n\t")
                                        _ = v16
                                    }
                                case "a":
                                    var v17 int32 = -262
                                    _ = v17
                                default:
                                    v18 := GenS2{F0: 0}
                                    _ = v18
                                }
                                if (fi13 > (fi9 % ((fi13) | 1))) {
                                    v19 := GenS2{F0: 1}
                                    _ = v19
                                }
                                var v20 uint64 = 456
                                v12 = append(v12, (fi13 + (fi9 & fi13)))
                                var v21 int64 = v14
                                if len(v12) > 0 {
                                    v12[0] = 316
                                }
                                _ = fi13
                                _ = v14
                                _ = v20
                                _ = v21
                            }
                            _ = v12
                        }
                    }
                    _ = fi9
                }
                v1 -= (-(v1))
                _ = v8
            }
            fmt.Println(v6, v6, v2)
            v22 := []int{32767, -648}
            _ = v22
        }
        v23 := GenS1{F0: v1, F1: -473}
        _ = v23
        _ = v6
    } else {
        v24 := v3[0]
        _ = v24
    }
    v25 := v3[139]
    v26 := GenS3{F0: 2147483647, F1: uint(v2), F2: 2147483647}
    v25 += (v25 - v25)
    v27 := GenS4{F0: -365, F1: 32767}
    v28 := []int{128}
    fmt.Println(v5)
    if false {
        v2 -= v2
        var v29 int = (v25 | -2147483648)
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
