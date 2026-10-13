package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16 }

func main() {
    v1 := []uint32{}
    var v2 uint16 = 104
    v3 := make(map[int]bool)
    v4 := GenS1{F0: 154}
    var v5 int = 0
    var v6 int16 = 1
    v2 = v2
    for fi7 := 0; fi7 < 1; fi7++ {
        v6 = v6
        if (42 >= 9223372036854775807) {
            var v8 uint8 = 128
            fmt.Println(fi7, v8)
            _ = v8
        } else {
            if true {
                v9 := GenS1{F0: 42}
                _ = v9
            } else {
                v10 := GenS2{F0: v6}
                var v11 uint64 = 1
                _ = v10
                _ = v11
            }
            v2 -= (^(v2))
            if true {
                v5 -= (fi7 / ((fi7) | 1))
            }
        }
        _ = fi7
    }
    fmt.Println(v6)
    var v12 uint8 = 58
    v5 -= (v5 ^ -406)
    v1 = append(v1, 1)
    v2 += (v2 + v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
