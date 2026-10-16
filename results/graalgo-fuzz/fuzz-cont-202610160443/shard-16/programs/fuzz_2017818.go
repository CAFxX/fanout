package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 uint32; F1 int32; F2 int8 }

func main() {
    v1 := make(map[string]int)
    v2 := []string{"\n\t", "a", "hello", "hello"}
    var v3 int8 = -1
    var v4 bool = (((1.5 + 3.14159) != 0.1) || false)
    v5 := make(map[uint32]int)
    v6 := make(map[string]uint32)
    if len(v2) > 0 {
        v2[0] = "a"
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v7 float32 = -2.25
    v4 = v4
    v8 := GenS1{F0: 4294967295, F1: int(v3)}
    v9 := GenS1{F0: 36, F1: -9223372036854775808}
    fmt.Println(v4, v7, v7)
    for fi10 := 0; fi10 < 8; fi10++ {
        for fi11 := 0; fi11 < 17; fi11++ {
            var v12 int = (^(fi10))
            _ = fi11
            _ = v12
        }
        fmt.Println(v4, v7, fi10)
        _ = fi10
    }
    v13 := GenS1{F0: 1, F1: 0}
    v14 := GenS2{F0: 2, F1: 42, F2: v3}
    var v15 int = 42
    for fi16 := 0; fi16 < 12; fi16++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi16
    }
    fmt.Println(v4, v3, v7)
    var v17 float64 = (-2.25 * 1.5)
    v7 = (((3.14159 * float32((v17 + (v17 / ((v17 * (v17 * (3.14159 - v17))) + ((1.5 * (-2.25 * 0.1)) / float64(v7))))))) * v7) + ((v7 * (v7 - v7)) - float32(v17)))
    fmt.Println(v7, v15, v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
