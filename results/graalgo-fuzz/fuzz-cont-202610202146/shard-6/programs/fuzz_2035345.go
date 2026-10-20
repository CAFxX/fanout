package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    var v1 uint8 = 167
    v2 := make(map[uint32]int)
    var v3 float64 = ((0.1 * ((0.0 - (0.0 / (0.1 / 1e10))) / 0.1)) + 1e10)
    v4 := make(map[uint32]string)
    var v5 float64 = v3
    v6 := make(map[uint32]int)
    var v7 uint8 = 242
    v8 := GenS1{F0: "\n\t"}
    v1 = 205
    fmt.Println(v5, v1)
    v6[905] = 0
    for fi9 := 0; fi9 < 15; fi9++ {
        var v10 int32 = -1
        for fi11 := 0; fi11 < 20; fi11++ {
            v10 -= (v10 - v10)
            v12 := []int{1}
            _ = fi11
            _ = v12
        }
        _ = fi9
        _ = v10
    }
    var v13 bool = ((1 == 64) || false)
    fmt.Println(v3)
    v7 += v1
    fmt.Println(v5, v1)
    v2[821] = 9223372036854775807
    var v14 int64 = 9223372036854775807
    v5 += v5
    var v15 uint = 845
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
