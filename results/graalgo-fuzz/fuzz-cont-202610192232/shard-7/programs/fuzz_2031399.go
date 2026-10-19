package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 uint }
type GenS2 struct { F0 uint16; F1 uint8 }
type GenS3 struct { F0 string; F1 int }

func main() {
    v1 := []int{32767}
    v2 := make(map[int]int)
    var v3 float32 = 1e10
    var v4 uint16 = 32767
    if (v4 == 65535) {
        v5 := GenS1{F0: 0, F1: (v4 * (v4 + v4)), F2: 42}
        _ = v5
    }
    fmt.Println(v3)
    v4 -= (v4 + v4)
    v4 -= v4
    v6 := GenS2{F0: v4, F1: 128}
    if (64 >= 128) {
        _ = len(v1)
        for fi7 := 0; fi7 < 6; fi7++ {
            fmt.Println(v4, v3, fi7)
            _ = fi7
        }
        if (!false) {
            for fi8 := 0; fi8 < 15; fi8++ {
                if true {
                    v4 -= 2
                    v4 += v4
                }
                if false {
                    if true {
                        var v9 bool = ((762 == -1) && false)
                        var v10 uint64 = 877
                        v11 := GenS2{F0: v4, F1: 255}
                        _ = v10
                        _ = v11
                        _ = v9
                    } else {
                        for fi12 := 0; fi12 < 10; fi12++ {
                            var v13 string = "a"
                            fmt.Println(fi12, fi8, v13)
                            _ = fi12
                            _ = v13
                        }
                        fmt.Println(v3, v4, v4)
                    }
                    var v14 uint16 = (^(v4))
                    v15 := v2[fi8]
                    _ = v14
                    _ = v15
                } else {
                    v16 := GenS3{F0: "\n\t", F1: fi8}
                    var v17 int32 = 417
                    var v18 float32 = float32(float64(-2.25))
                    _ = v16
                    _ = v17
                    _ = v18
                }
                _ = fi8
            }
        }
    } else {
        if true {
            var v19 uint32 = 4294967295
            fmt.Println(v3, v19, v19)
            _ = v19
        }
    }
    for fi20 := 0; fi20 < 3; fi20++ {
        var v21 int64 = 42
        var v22 float32 = 0.1
        _ = fi20
        _ = v21
        _ = v22
    }
    v4 = (v4 | (v4 - (v4 % (((^(v4))) | 1))))
    if len(v1) > 0 {
        v1[0] = 9223372036854775807
    }
    if false {
        v4 = v4
        if len(v1) > 0 {
            v1[0] = 0
        }
        v3 = v3
    } else {
        v4 = v4
        if true {
            v3 -= v3
        } else {
            var v23 uint32 = 65535
            v4 = v4
            _ = v23
        }
    }
    var v24 int = 138
    var v25 uint8 = 2
    v4 = 255
    fmt.Println(v3, v4, v4)
    v25 = (v25 + (v25 - v25))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v6
}
