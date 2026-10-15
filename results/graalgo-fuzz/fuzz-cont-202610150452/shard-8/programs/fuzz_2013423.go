package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 uint32 }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]int)
    v3 := make(map[int]bool)
    v4 := make(map[string]string)
    v5 := []int{32768}
    v6 := make(map[string]int)
    var v7 float32 = 1.5
    fmt.Println(v7)
    var v8 int8 = -59
    if len(v5) > 0 {
        v5[0] = -294
    }
    for fi9 := 0; fi9 < 12; fi9++ {
        for fi10 := 0; fi10 < 20; fi10++ {
            var v11 uint8 = 255
            _ = fi10
            _ = v11
        }
        for fi12 := 0; fi12 < 7; fi12++ {
            var v13 uint16 = 65535
            v2[fi12] = fi9
            _ = fi12
            _ = v13
        }
        _ = fi9
    }
    for fi14 := 0; fi14 < 2; fi14++ {
        for fi15 := 0; fi15 < 9; fi15++ {
            var v16 int = (fi15 + 302)
            var v17 bool = false
            _ = fi15
            _ = v16
            _ = v17
        }
        fmt.Println(v7, v8, fi14)
        _ = fi14
    }
    v18 := v2[908]
    v19 := GenS1{F0: 2, F1: 1, F2: 2147483648}
    var v20 bool = (!(710 > 42))
    fmt.Println(v18)
    v21 := GenS1{F0: 0, F1: 437, F2: 2147483647}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
