package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int64; F1 int8; F2 uint16 }
type GenS5 struct { F0 int8 }

func main() {
    v1 := []uint32{}
    v2 := make(map[int]uint32)
    var v3 bool = (1000 > 32767)
    var v4 float32 = (1.5 + 0.1)
    v5 := []int{-1, 0, -1, 127}
    var v6 int32 = 0
    if v3 {
        fmt.Println(v3, v3)
        for fi7 := 0; fi7 < 2; fi7++ {
            _ = len(v1)
            _ = fi7
        }
        if (false || (!false)) {
            v8 := GenS1{F0: 0}
            _ = v8
        } else {
            v9 := GenS1{F0: 611}
            if v3 {
                v10 := GenS2{F0: ("a" + "\n\t")}
                v11 := GenS1{F0: 2}
                _ = v10
                _ = v11
            }
            if v3 {
                v6 -= 2147483647
                v6 = 702
            }
            _ = v9
        }
    }
    delete(v2, 874)
    fmt.Println(v3, v6)
    fmt.Println(v6, v4, v6)
    for fi12 := 0; fi12 < 16; fi12++ {
        fmt.Println(v6, v6)
        _ = fi12
    }
    var v13 int = 4294967295
    v14 := GenS3{F0: 128}
    var v15 float64 = (1.5 + 0.1)
    var v16 int64 = -1
    v4 += -2.25
    fmt.Println(v15, v15)
    for fi17 := 0; fi17 < 8; fi17++ {
        v2[(fi17 * v13)] = 942
        if (0 < 18446744073709551615) {
            if len(v5) > 0 {
                v5[0] = fi17
            }
        } else {
            var v18 int16 = 32767
            for fi19 := 0; fi19 < 5; fi19++ {
                var v20 int16 = (v18 * v18)
                if v3 {
                    var v21 uint32 = 999
                    switch (v13 / (((^(fi19))) | 1)) {
                    case 4294967295:
                        if true {
                            v22 := GenS2{F0: "foo bar"}
                            _ = v22
                        }
                    case 1:
                        v18 = int16(v21)
                    default:
                        v13 += (v13 + 415)
                    }
                    v23 := v2[v13]
                    _ = v21
                    _ = v23
                } else {
                    fmt.Println(v13)
                    v24 := GenS1{F0: 1}
                    if (120 != 42) {
                        switch (fi19 & 9223372036854775807) {
                        case 32767:
                            delete(v2, fi17)
                        case -2147483648:
                            v15 -= (v15 * v15)
                        default:
                            var v25 float64 = v15
                            _ = v25
                        }
                        fmt.Println(v16)
                        var v26 bool = true
                        _ = v26
                    } else {
                        v27 := GenS3{F0: (fi17 % (((-(v13))) | 1))}
                        _ = v27
                    }
                    _ = v24
                }
                var v28 uint16 = 32767
                _ = fi19
                _ = v20
                _ = v28
            }
            v29 := GenS4{F0: v16, F1: -128, F2: 0}
            _ = v18
            _ = v29
        }
        var v30 uint8 = 238
        _ = fi17
        _ = v30
    }
    if false {
        v31 := GenS5{F0: -128}
        var v32 int64 = (v16 + v16)
        _ = v31
        _ = v32
    }
    _ = len(v2)
    v13 += 4294967295
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
