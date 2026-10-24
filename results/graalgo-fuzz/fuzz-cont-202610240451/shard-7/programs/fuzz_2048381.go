package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 int32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int16 }

func main() {
    var v1 int64 = 0
    var v2 uint = 18446744073709551615
    v3 := make(map[string]bool)
    var v4 uint64 = 219
    var v5 uint64 = v4
    for fi6 := 0; fi6 < 6; fi6++ {
        v7 := GenS1{F0: 2, F1: (fi6 % ((int(v1)) | 1)), F2: int32(v1)}
        if ((false && (-1 <= -714)) && false) {
            for fi8 := 0; fi8 < 13; fi8++ {
                v9 := GenS2{F0: (255 <= 2)}
                _ = fi8
                _ = v9
            }
        }
        _ = fi6
        _ = v7
    }
    var v10 uint32 = 8
    v11 := []int{-887, 9223372036854775807, 0, 32767}
    var v12 uint64 = 2
    v12 = v5
    v13 := GenS1{F0: 1, F1: 32768, F2: int32(v5)}
    _ = len(v3)
    for fi14 := 0; fi14 < 9; fi14++ {
        if true {
            v15 := GenS3{F0: 1000}
            _ = v15
        } else {
            fmt.Println(v12, v10, v2)
            var v16 int = (-(fi14))
            _ = v16
        }
        v17 := v3["foo bar"]
        if (1 == v10) {
            fmt.Println(fi14)
            if (("hello" + "") > "foo bar") {
                fmt.Println(v10, v10)
                fmt.Println(v4, v4)
                if v17 {
                    switch fi14 {
                    case 127:
                        fmt.Println(v10, v5, v17)
                    case 4294967295:
                        var v18 uint64 = 0
                        _ = v18
                    case 2147483647:
                        var v19 uint16 = 1
                        _ = v19
                    default:
                        for fi20 := 0; fi20 < 19; fi20++ {
                            var v21 float64 = 1e10
                            var v22 string = "\n\t"
                            fmt.Println(v2, v10)
                            _ = fi20
                            _ = v21
                            _ = v22
                        }
                    }
                    v1 = (v1 - (v1 - (v1 + 42)))
                    var v23 uint16 = 799
                    _ = v23
                } else {
                    for fi24 := 0; fi24 < 13; fi24++ {
                        v17 = (v2 != (v2 % (((v2 / ((v2) | 1))) | 1)))
                        _ = fi24
                    }
                    v2 -= (v2 | v2)
                    v25 := GenS3{F0: -213}
                    _ = v25
                }
            } else {
                v12 = v5
                v26 := GenS2{F0: true}
                for fi27 := 0; fi27 < 13; fi27++ {
                    var v28 int = fi27
                    if len(v11) > 0 {
                        v11[0] = (v28 * -157)
                    }
                    fmt.Println(fi14, v5)
                    _ = fi27
                    _ = v28
                }
                _ = v26
            }
        }
        _ = fi14
        _ = v17
    }
    if true {
        if len(v11) > 0 {
            v11[0] = 1
        }
    }
    v29 := GenS1{F0: 207, F1: 32768, F2: 2147483647}
    fmt.Println(v1, v2, v5)
    v3[("a" + "foo bar")] = false
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
