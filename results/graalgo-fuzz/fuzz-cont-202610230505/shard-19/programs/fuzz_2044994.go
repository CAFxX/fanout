package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16 }

func main() {
    var v1 uint64 = 489
    v2 := []int{4294967295}
    var v3 int = 32768
    v4 := []uint32{}
    v5 := make(map[int]bool)
    v1 = 2
    fmt.Println(v1, v1)
    v6 := v5[v3]
    for fi7 := 0; fi7 < 3; fi7++ {
        _ = len(v5)
        _ = len(v4)
        v8 := GenS1{F0: 2147483647}
        _ = fi7
        _ = v8
    }
    v6 = false
    v9 := GenS1{F0: 106}
    if len(v2) > 0 {
        v2[0] = v3
    }
    v6 = v6
    v10 := GenS2{F0: 0}
    v1 = (v1 >> ((65535) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
