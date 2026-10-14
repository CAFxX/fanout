package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int8; F1 uint64; F2 bool }

func main() {
    var v1 uint32 = 4294967295
    v2 := make(map[string]string)
    v3 := make(map[int]bool)
    v4 := []int{0}
    var v5 uint32 = (^(v1))
    var v6 int32 = 585
    var v7 uint64 = 2
    v8 := GenS1{F0: 100}
    v9 := GenS2{F0: 2}
    var v10 float64 = -2.25
    v7 -= 42
    v7 += 1
    if len(v4) > 0 {
        v4[0] = -878
    }
    var v11 int32 = (v6 / (((v6 % (((v6 / ((v6) | 1))) | 1))) | 1))
    v1 += v1
    v1 = (v5 >> ((18446744073709551615) & 63))
    fmt.Println(v10, v11)
    v12 := GenS3{F0: 0, F1: v7, F2: (!false)}
    v13 := GenS2{F0: 2}
    for fi14 := 0; fi14 < 11; fi14++ {
        v6 = (v6 | v11)
        _ = fi14
    }
    v2["hello"] = "a"
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
