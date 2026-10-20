package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8; F1 uint8 }
type GenS3 struct { F0 int64; F1 bool }

func main() {
    var v1 int32 = 42
    var v2 bool = false
    var v3 uint64 = 0
    v2 = v2
    if v2 {
        fmt.Println(v3, v1)
        if v2 {
            if v2 {
                for fi4 := 0; fi4 < 16; fi4++ {
                    v1 = v1
                    v5 := GenS1{F0: -762}
                    var v6 uint16 = 917
                    _ = fi4
                    _ = v5
                    _ = v6
                }
            } else {
                for fi7 := 0; fi7 < 8; fi7++ {
                    switch fi7 {
                    case 284:
                        for fi8 := 0; fi8 < 2; fi8++ {
                            var v9 int64 = 9223372036854775807
                            _ = fi8
                            _ = v9
                        }
                    case -9223372036854775808:
                        for fi10 := 0; fi10 < 16; fi10++ {
                            v3 = v3
                            _ = fi10
                        }
                    }
                    _ = fi7
                }
                v11 := GenS2{F0: -128, F1: 1}
                if (!((255 == 0) && (0 > 128))) {
                    if v2 {
                        var v12 int8 = 0
                        _ = v12
                    }
                    for fi13 := 0; fi13 < 13; fi13++ {
                        v3 = v3
                        v14 := GenS2{F0: 1, F1: 42}
                        _ = fi13
                        _ = v14
                    }
                }
                _ = v11
            }
            v15 := []int{-225, 1}
            v2 = (3.14159 > (float32((1.5 - 0.1)) * (3.14159 * float32(-2.25))))
            _ = v15
        }
        if v2 {
            v1 -= (v1 >> ((905) & 63))
        } else {
            for fi16 := 0; fi16 < 1; fi16++ {
                v3 -= (v3 ^ (v3 + v3))
                v17 := []int{fi16, fi16}
                _ = fi16
                _ = v17
            }
            fmt.Println(v1, v1, v2)
            v3 = v3
        }
    } else {
        v18 := []int{32767, 42, 0}
        _ = v18
    }
    switch 0 {
    case -2147483648:
        v1 = (v1 / ((933) | 1))
    case 517:
        var v19 int32 = (v1 | v1)
        _ = v19
    }
    switch "foo bar" {
    case "":
        for fi20 := 0; fi20 < 6; fi20++ {
            for fi21 := 0; fi21 < 13; fi21++ {
                v1 -= (v1 - (-(v1)))
                v22 := []int{466, (fi20 & int(v1)), fi20, 0}
                _ = fi21
                _ = v22
            }
            _ = fi20
        }
    }
    v23 := []int{32768}
    v23 = append(v23, -9223372036854775808)
    for fi24 := 0; fi24 < 5; fi24++ {
        if len(v23) > 0 {
            v23[0] = (fi24 | 9223372036854775807)
        }
        _ = fi24
    }
    v25 := GenS3{F0: 879, F1: true}
    v2 = v2
    if len(v23) > 0 {
        v23[0] = int(v3)
    }
    var v26 bool = false
    v27 := GenS1{F0: 2147483647}
    var v28 uint64 = 325
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v23
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
}
