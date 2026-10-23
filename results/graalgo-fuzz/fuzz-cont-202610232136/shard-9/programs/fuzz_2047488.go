package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32 }
type GenS2 struct { F0 bool; F1 uint64; F2 uint64 }
type GenS3 struct { F0 int8; F1 uint }
func genFn0(p0 int) uint32 {
    v5 := []int{}
    fmt.Println(p0)
    var v6 int16 = -1
    for fi7 := 0; fi7 < 15; fi7++ {
        p0 -= -550
        var v8 float32 = (((0.0 / -2.25) * (0.0 - 0.0)) * 0.1)
        _ = fi7
        _ = v8
    }
    v6 -= v6
    _ = p0
    _ = v5
    _ = v6
    return 2147483648
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[string]string)
    var v3 int16 = 32767
    v4 := []int{-127}
    var v9 int32 = 1
    v10 := GenS1{F0: 127, F1: v9}
    if false {
        _ = len(v4)
        if false {
            if (-716 >= -9223372036854775808) {
                v11 := v2["foo bar"]
                _ = v11
            } else {
                v12 := GenS1{F0: 127, F1: 753}
                _ = v12
            }
        } else {
            if false {
                for fi13 := 0; fi13 < 13; fi13++ {
                    if false {
                        v9 += v9
                        if (!true) {
                            fmt.Println(v3)
                            for fi14 := 0; fi14 < 3; fi14++ {
                                var v15 uint16 = 304
                                var v16 int16 = (v3 << ((42) & 63))
                                v1[264] = 32767
                                _ = fi14
                                _ = v15
                                _ = v16
                            }
                            for fi17 := 0; fi17 < 9; fi17++ {
                                v18 := GenS1{F0: 42, F1: 0}
                                fmt.Println(v9)
                                _ = fi17
                                _ = v18
                            }
                        }
                    }
                    v3 -= v3
                    _ = fi13
                }
                v19 := GenS2{F0: (!false), F1: 55, F2: 0}
                _ = v19
            } else {
                v9 = (v9 % (((v9 ^ -728)) | 1))
            }
        }
    }
    for fi20 := 0; fi20 < 8; fi20++ {
        v21 := GenS3{F0: -1, F1: 977}
        _ = fi20
        _ = v21
    }
    v3 += -808
    var v22 float32 = float32((float64(3.14159) / (-2.25 + 0.1)))
    var v23 uint32 = 2
    v24 := GenS1{F0: 11, F1: v9}
    v25 := GenS3{F0: 42, F1: 42}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
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
