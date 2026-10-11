package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 uint8 }
type GenS2 struct { F0 uint16 }

func main() {
    v1 := []int{-1}
    var v2 bool = (188 >= 0)
    v3 := make(map[string]string)
    v2 = (!(!v2))
    if v2 {
        for fi4 := 0; fi4 < 13; fi4++ {
            var v5 bool = false
            for fi6 := 0; fi6 < 7; fi6++ {
                var v7 int16 = -219
                if (!(!v2)) {
                    var v8 int = fi4
                    _ = v8
                }
                _ = fi6
                _ = v7
            }
            _ = fi4
            _ = v5
        }
    }
    v2 = (0 > 32767)
    var v9 uint32 = 2
    if (255 > 1) {
        var v10 int = 9223372036854775807
        v11 := GenS1{F0: 0, F1: 0, F2: 2}
        _ = v10
        _ = v11
    } else {
        v1 = append(v1, -1)
        fmt.Println(v9, v9)
        v12 := GenS2{F0: 32767}
        _ = v12
    }
    if v2 {
        var v13 uint64 = 2
        _ = len(v1)
        if v2 {
            v13 -= v13
            _ = len(v1)
            fmt.Println(v2)
        } else {
            v13 = 189
            var v14 int = 2147483647
            _ = v14
        }
        _ = v13
    }
    for fi15 := 0; fi15 < 3; fi15++ {
        v16 := GenS2{F0: 42}
        var v17 int = (fi15 + (-(fi15)))
        _ = fi15
        _ = v16
        _ = v17
    }
    fmt.Println(v2, v9)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v9
}
