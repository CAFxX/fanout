package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 int64; F1 int }
type GenS3 struct { F0 uint64; F1 int8; F2 int64 }

func main() {
    var v1 float32 = (((3.14159 / ((1e10 - 0.1) + float32(1e10))) / 1.0) + float32(3.14159))
    var v2 bool = false
    v3 := make(map[int]bool)
    var v4 uint = 365
    v5 := make(map[int]bool)
    v6 := make(map[uint32]string)
    var v7 float32 = v1
    if v2 {
        if v2 {
            var v8 int8 = -1
            v9 := GenS1{F0: 54, F1: ("" + (("a" + "a") + "a"))}
            _ = v8
            _ = v9
        }
    }
    var v10 string = "\n\t"
    for fi11 := 0; fi11 < 17; fi11++ {
        v10 = (v10 + v10)
        _ = fi11
    }
    var v12 float32 = (v1 - 1.5)
    var v13 int = 505
    var v14 uint64 = 18446744073709551615
    v15 := []int{1, (v13 / (((^(v13))) | 1))}
    var v16 string = v10
    var v17 uint8 = 43
    if v2 {
        switch "a" {
        case "":
            for fi18 := 0; fi18 < 13; fi18++ {
                var v19 uint8 = (v17 - v17)
                for fi20 := 0; fi20 < 9; fi20++ {
                    var v21 uint8 = v17
                    if (v2 || v2) {
                        fmt.Println(fi20)
                    } else {
                        if v2 {
                            var v22 float32 = (v7 - (0.0 * (float32(((((0.1 + 0.1) * -2.25) - -2.25) - 1.5)) * -2.25)))
                            _ = v22
                        } else {
                            if (-359 != 1) {
                                if true {
                                    var v23 bool = v2
                                    var v24 int8 = 100
                                    _ = v23
                                    _ = v24
                                }
                                var v25 int64 = 669
                                fmt.Println(v25)
                                _ = v25
                            }
                            v6[0] = (("foo bar" + "hello") + v10)
                        }
                        v12 += -2.25
                    }
                    _ = fi20
                    _ = v21
                }
                v26 := GenS2{F0: -9223372036854775808, F1: (fi18 - (fi18 % (((fi18 << ((uint(v14)) & 63))) | 1)))}
                _ = fi18
                _ = v19
                _ = v26
            }
        default:
            for fi27 := 0; fi27 < 9; fi27++ {
                v28 := GenS3{F0: 166, F1: 62, F2: 940}
                v6[517] = v10
                _ = fi27
                _ = v28
            }
        }
        for fi29 := 0; fi29 < 9; fi29++ {
            v30 := v3[-406]
            v13 += (fi29 | v13)
            fmt.Println(v1)
            _ = fi29
            _ = v30
        }
    } else {
        v14 = (v14 ^ uint64(v4))
    }
    fmt.Println(v4, v4)
    var v31 int16 = -298
    v7 = 0.0
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
