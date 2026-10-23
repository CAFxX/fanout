package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0() uint16 {
    var v7 uint32 = 2
    v8 := []int{-403, 720}
    switch "" {
    case "foo bar":
        v7 = (v7 + (v7 % (((v7 / (((v7 * (v7 | (v7 & v7)))) | 1))) | 1)))
    }
    v7 = v7
    _ = v7
    _ = v8
    return 42
}

func main() {
    var v1 int = 2147483647
    v2 := make(map[string]int)
    v3 := make(map[int]uint32)
    var v4 float32 = (1.5 + ((1.5 * -2.25) / 3.14159))
    var v5 int16 = 915
    v6 := make(map[uint32]int)
    v4 = v4
    v9 := []int{v1, v1}
    var v10 uint16 = genFn0()
    v11 := GenS1{F0: 42}
    v10 += v10
    v10 = (v10 & genFn0())
    for fi12 := 0; fi12 < 20; fi12++ {
        var v13 uint8 = 255
        _ = fi12
        _ = v13
    }
    var v14 int64 = 455
    fmt.Println(v14)
    _ = len(v9)
    v9 = append(v9, v1)
    v5 = (v5 << ((2) & 63))
    _ = len(v2)
    var v15 int8 = -1
    v16 := GenS1{F0: -2147483648}
    v15 -= v15
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
