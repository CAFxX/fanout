package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[int]int)
    var v3 int32 = -1
    for fi4 := 0; fi4 < 16; fi4++ {
        if (false || true) {
            for fi5 := 0; fi5 < 5; fi5++ {
                _ = len(v2)
                _ = fi5
            }
            for fi6 := 0; fi6 < 16; fi6++ {
                v3 += 0
                v7 := GenS1{F0: 18446744073709551615, F1: 754}
                fmt.Println(fi4, fi6)
                _ = fi6
                _ = v7
            }
            v8 := GenS1{F0: 784, F1: 1}
            _ = v8
        } else {
            var v9 bool = (true && false)
            _ = v9
        }
        _ = fi4
    }
    for fi10 := 0; fi10 < 1; fi10++ {
        v3 = -1
        var v11 uint = 18446744073709551615
        v1[fi10] = 2147483647
        _ = fi10
        _ = v11
    }
    var v12 int64 = 1
    fmt.Println(v12, v12)
    v12 -= (-(v12))
    fmt.Println(v3)
    v12 = v12
    v13 := v2[0]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
}
