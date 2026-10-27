package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 int8 }

func main() {
    var v1 int16 = 42
    var v2 float32 = 0.1
    var v3 uint8 = 195
    v4 := make(map[int]int)
    v5 := GenS1{F0: -421, F1: 42}
    if true {
        if false {
            if (!true) {
                switch 301 {
                case -9223372036854775808:
                    v6 := v4[123]
                    _ = v6
                case 127:
                    var v7 int64 = -9223372036854775808
                    _ = v7
                default:
                    if false {
                        v2 += float32(((float64(v2) + (3.14159 / ((((-2.25 + 1e10) / -2.25) - 1.5) * -2.25))) - (1.5 / 0.1)))
                        fmt.Println(v3, v3, v1)
                        var v8 uint64 = 4294967295
                        _ = v8
                    } else {
                        var v9 bool = (!false)
                        if (42 != -61) {
                            v10 := []int{42, -738, -860, -918}
                            for fi11 := 0; fi11 < 11; fi11++ {
                                for fi12 := 0; fi12 < 2; fi12++ {
                                    var v13 uint8 = (v3 % ((42) | 1))
                                    _ = fi12
                                    _ = v13
                                }
                                if v9 {
                                    var v14 int32 = -325
                                    v1 -= -32768
                                    _ = v14
                                }
                                _ = len(v4)
                                _ = fi11
                            }
                            v15 := GenS2{F0: 1}
                            _ = v10
                            _ = v15
                        }
                        for fi16 := 0; fi16 < 12; fi16++ {
                            var v17 uint = 705
                            v18 := []int{32768, (fi16 >> ((v17) & 63)), (fi16 >> (((v17 * (^(v17)))) & 63)), fi16}
                            _ = fi16
                            _ = v17
                            _ = v18
                        }
                        _ = v9
                    }
                }
            } else {
                fmt.Println(v1, v3, v3)
            }
            var v19 float32 = 0.0
            v19 += float32(1e10)
            _ = v19
        }
        v20 := GenS1{F0: 64, F1: int8(v1)}
        _ = v20
    }
    fmt.Println(v3, v3)
    for fi21 := 0; fi21 < 16; fi21++ {
        for fi22 := 0; fi22 < 14; fi22++ {
            v2 -= v2
            var v23 int32 = 798
            var v24 float64 = 1e10
            _ = fi22
            _ = v23
            _ = v24
        }
        v25 := []int{(fi21 & fi21), (-(fi21)), (-(fi21)), int(v3)}
        _ = fi21
        _ = v25
    }
    switch 423 {
    case 4294967295:
        v1 = v1
    }
    if (!false) {
        v26 := GenS1{F0: 9223372036854775807, F1: 1}
        _ = v26
    }
    var v27 uint16 = 42
    fmt.Println(v1)
    switch "\n\t" {
    case "hello":
        v3 -= v3
    default:
        v28 := GenS2{F0: -1}
        _ = v28
    }
    var v29 int16 = (v1 ^ v1)
    if true {
        for fi30 := 0; fi30 < 10; fi30++ {
            v31 := []int{int(v29), fi30, fi30, fi30}
            v31 = append(v31, fi30)
            _ = fi30
            _ = v31
        }
        if false {
            var v32 int = 9223372036854775807
            var v33 float32 = v2
            _ = v32
            _ = v33
        } else {
            v3 = v3
            v34 := GenS1{F0: -758, F1: 42}
            _ = v34
        }
    }
    if (127 <= 0) {
        v1 = 32767
    } else {
        v3 = (v3 + (v3 + v3))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v27
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
