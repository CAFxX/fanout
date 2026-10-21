package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 uint64 }

func main() {
    v1 := make(map[string]int)
    var v2 int8 = 0
    var v3 int16 = 32767
    var v4 string = "a"
    var v5 uint64 = 1
    v3 -= (v3 >> ((1) & 63))
    fmt.Println(v4, v3)
    v6 := []int{32768, 42, -748}
    fmt.Println(v4, v4, v5)
    for fi7 := 0; fi7 < 15; fi7++ {
        var v8 uint32 = 920
        _ = fi7
        _ = v8
    }
    var v9 bool = (v2 >= (v2 % (((v2 >> ((42) & 63))) | 1)))
    _ = len(v6)
    v6 = append(v6, 947)
    var v10 uint16 = 705
    v4 = ""
    v4 = v4
    v11 := GenS1{F0: v3, F1: 0, F2: v5}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
