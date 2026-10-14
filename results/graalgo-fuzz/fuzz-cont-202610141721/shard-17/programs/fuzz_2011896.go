package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 uint16 }
type GenS2 struct { F0 int }
func genFn0(p0 string, p1 uint) uint64 {
    v6 := GenS1{F0: 182, F1: 442, F2: 1}
    p0 = (p0 + "foo bar")
    v7 := GenS2{F0: 1}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 50
}

func main() {
    v1 := []int{-1, 42, 9223372036854775807}
    var v2 int32 = 943
    v3 := make(map[int]uint32)
    var v4 float32 = 3.14159
    var v5 uint = 655
    fmt.Println(v4, v4, v5)
    v8 := GenS2{F0: -9223372036854775808}
    _ = len(v1)
    v9 := GenS2{F0: 32768}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v10 := v3[-9223372036854775808]
    v4 = float32((1.5 - 1.5))
    var v11 uint8 = 95
    for fi12 := 0; fi12 < 9; fi12++ {
        v10 += v10
        _ = fi12
    }
    v11 = 2
    v13 := GenS2{F0: 128}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
