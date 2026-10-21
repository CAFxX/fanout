package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 int32 }
type GenS2 struct { F0 uint; F1 int64; F2 bool }
type GenS3 struct { F0 uint }

func main() {
    var v1 float64 = ((0.1 * (float64((1.5 / (1.5 + -2.25))) * (0.1 * 3.14159))) + (-2.25 * (3.14159 - (1e10 + 1e10))))
    var v2 int16 = 0
    var v3 uint = 1
    v4 := GenS1{F0: 1, F1: v2, F2: 1}
    v5 := GenS1{F0: 783, F1: -686, F2: -1}
    v6 := GenS1{F0: 0, F1: v2, F2: 0}
    v3 -= 924
    if (0.0 == float32((1.5 + (0.1 + float64(float32((v1 + v1))))))) {
        v7 := GenS2{F0: v3, F1: 9223372036854775807, F2: false}
        for fi8 := 0; fi8 < 6; fi8++ {
            for fi9 := 0; fi9 < 10; fi9++ {
                var v10 uint = (v3 ^ 1)
                _ = fi9
                _ = v10
            }
            _ = fi8
        }
        _ = v7
    } else {
        fmt.Println(v1)
        fmt.Println(v2, v1)
        for fi11 := 0; fi11 < 14; fi11++ {
            v1 = (v1 / (-2.25 / ((float64(0.1) / 3.14159) * float64(0.1))))
            if ((!(("" + ("" + "\n\t")) >= "a")) && (uint16(v3) != 42)) {
                switch fi11 {
                case 32768:
                    v12 := GenS1{F0: 42, F1: (v2 >> ((42) & 63)), F2: 941}
                    _ = v12
                case -2:
                    var v13 int64 = 9223372036854775807
                    _ = v13
                default:
                    if false {
                        fmt.Println(v1, v3, v1)
                        if (1 != v2) {
                            var v14 string = ("foo bar" + "foo bar")
                            for fi15 := 0; fi15 < 10; fi15++ {
                                v1 = (v1 + 3.14159)
                                var v16 int8 = 11
                                var v17 uint = 65535
                                _ = fi15
                                _ = v16
                                _ = v17
                            }
                            v2 = v2
                            _ = v14
                        }
                        if false {
                            if true {
                                v3 = v3
                            } else {
                                var v18 float64 = 1e10
                                v19 := []int{(fi11 % ((-939) | 1)), fi11, fi11}
                                _ = len(v19)
                                _ = v18
                                _ = v19
                            }
                        } else {
                            var v20 uint64 = 4294967295
                            _ = v20
                        }
                    }
                }
                v21 := []int{}
                _ = v21
            }
            _ = fi11
        }
    }
    var v22 bool = false
    for fi23 := 0; fi23 < 1; fi23++ {
        for fi24 := 0; fi24 < 20; fi24++ {
            var v25 int16 = v2
            _ = fi24
            _ = v25
        }
        _ = fi23
    }
    v26 := GenS3{F0: v3}
    v2 += (v2 * v2)
    fmt.Println(v1, v1)
    if (v22 || (!(65535 != 4294967295))) {
        var v27 int8 = -1
        v1 = -2.25
        _ = v27
    } else {
        if v22 {
            fmt.Println(v22, v1, v1)
        }
    }
    v28 := GenS1{F0: 2147483647, F1: v2, F2: -580}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v22
    _ = v26
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
