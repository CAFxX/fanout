package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    var v1 uint64 = 2
    var v2 bool = (!false)
    v3 := make(map[uint32]string)
    var v4 uint8 = 2
    for fi5 := 0; fi5 < 7; fi5++ {
        switch "a" {
        case "hello":
            v6 := []int{(-(fi5)), (^(fi5)), 32767}
            _ = v6
        case "a":
            for fi7 := 0; fi7 < 7; fi7++ {
                v8 := GenS1{F0: 2}
                _ = fi7
                _ = v8
            }
        }
        v9 := v3[42]
        _ = len(v3)
        _ = fi5
        _ = v9
    }
    v10 := []int{-597, 9223372036854775807}
    v11 := GenS1{F0: v1}
    v2 = (!(42 != v1))
    var v12 int32 = 2147483647
    fmt.Println(v4, v1, v12)
    v1 -= v1
    var v13 uint8 = (v4 + (^(v4)))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
