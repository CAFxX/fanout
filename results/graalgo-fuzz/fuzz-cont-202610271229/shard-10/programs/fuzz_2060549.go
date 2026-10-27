package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64; F1 uint16; F2 uint }

func main() {
    var v1 bool = false
    v2 := make(map[int]bool)
    var v3 int32 = 2147483647
    v4 := make(map[int]uint32)
    var v5 uint64 = 0
    if v1 {
        v6 := []int{-9223372036854775808, 128, 9223372036854775807}
        _ = v6
    }
    if v1 {
        for fi7 := 0; fi7 < 2; fi7++ {
            v8 := GenS1{F0: 802}
            var v9 string = "a"
            _ = fi7
            _ = v8
            _ = v9
        }
    }
    if v1 {
        fmt.Println(v1)
    }
    fmt.Println(v1)
    v10 := GenS2{F0: v5, F1: 1, F2: 686}
    v1 = false
    fmt.Println(v5, v1)
    var v11 uint16 = 1
    var v12 uint32 = 0
    fmt.Println(v5)
    v1 = ((v5 >= v5) && (!(995 > 835)))
    var v13 uint32 = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
