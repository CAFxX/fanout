package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[string]string)
    v3 := make(map[uint32]string)
    v4 := v1[547]
    for fi5 := 0; fi5 < 20; fi5++ {
        if false {
            _ = len(v3)
            fmt.Println(fi5)
            if (2147483647 >= -1) {
                v6 := GenS1{F0: -1}
                v7 := GenS1{F0: 2147483647}
                _ = v6
                _ = v7
            }
        }
        _ = fi5
    }
    fmt.Println(v4)
    fmt.Println(v4)
    var v8 float64 = float64(-2.25)
    var v9 uint64 = 1
    var v10 uint32 = 14
    v11 := []int{(v4 << ((255) & 63)), (^(v4))}
    v4 -= (v4 & v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
