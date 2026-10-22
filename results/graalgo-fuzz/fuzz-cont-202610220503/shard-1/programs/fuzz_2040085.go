package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 bool }

func main() {
    var v1 int = 51
    var v2 int32 = 421
    v3 := make(map[int]int)
    v4 := []string{"", "", "foo bar"}
    for fi5 := 0; fi5 < 12; fi5++ {
        v6 := GenS1{F0: -1, F1: 2}
        v2 = v2
        if (-9223372036854775808 >= -1) {
            if false {
                _ = len(v3)
                v1 = (fi5 % (((v1 & v1)) | 1))
                v7 := GenS2{F0: false}
                _ = v7
            }
        }
        _ = fi5
        _ = v6
    }
    var v8 uint32 = 679
    var v9 uint8 = 2
    fmt.Println(v9, v9)
    var v10 int16 = 1
    var v11 float64 = -2.25
    v9 = (v9 * v9)
    for fi12 := 0; fi12 < 19; fi12++ {
        _ = len(v3)
        fmt.Println(v9)
        v13 := GenS1{F0: int64(v1), F1: 643}
        _ = fi12
        _ = v13
    }
    fmt.Println(v8, v11, v1)
    var v14 uint32 = (v8 * (v8 / ((0) | 1)))
    for fi15 := 0; fi15 < 6; fi15++ {
        v1 += -223
        _ = fi15
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
