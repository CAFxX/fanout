package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    v1 := make(map[uint32]int)
    var v2 float32 = 1e10
    var v3 uint16 = 1
    var v4 uint32 = 2147483647
    v5 := GenS1{F0: 128}
    for fi6 := 0; fi6 < 1; fi6++ {
        fmt.Println(fi6)
        _ = len(v1)
        v3 += (v3 >> ((4294967295) & 63))
        _ = fi6
    }
    for fi7 := 0; fi7 < 1; fi7++ {
        if (v2 == v2) {
            v2 = 1e10
            fmt.Println(v3, fi7, v3)
        }
        _ = fi7
    }
    if false {
        v8 := v1[v4]
        _ = v8
    }
    if false {
        v9 := []int{42}
        _ = v9
    }
    switch 743 {
    case 127:
        v2 = (v2 / (v2 - float32(-2.25)))
    }
    fmt.Println(v2)
    var v10 float64 = ((0.1 / (float64(1.5) + (0.0 - -2.25))) - 1e10)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
