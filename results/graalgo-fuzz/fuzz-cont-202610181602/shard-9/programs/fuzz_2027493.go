package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 uint32 }
type GenS2 struct { F0 string }

func main() {
    v1 := make(map[string]int)
    var v2 uint64 = 1
    v3 := make(map[int]int)
    fmt.Println(v2)
    fmt.Println(v2)
    var v4 int32 = -2147483648
    if (!false) {
        var v5 int8 = 25
        for fi6 := 0; fi6 < 4; fi6++ {
            v7 := GenS1{F0: 293, F1: 232, F2: 215}
            v5 += 127
            if false {
                var v8 bool = true
                v9 := GenS1{F0: 42, F1: fi6, F2: 15}
                _ = v8
                _ = v9
            }
            _ = fi6
            _ = v7
        }
        _ = v5
    }
    var v10 float32 = ((-2.25 - ((3.14159 * 1.5) + 1e10)) * (1.5 * ((1.5 - (1.5 - 1.5)) - 0.0)))
    v2 = 18446744073709551615
    v11 := GenS2{F0: ""}
    for fi12 := 0; fi12 < 9; fi12++ {
        v2 += (v2 & (v2 & v2))
        v13 := []int{(-(fi12)), int(v4), (fi12 << ((18446744073709551615) & 63)), fi12}
        _ = fi12
        _ = v13
    }
    var v14 uint64 = 0
    v10 = 0.1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
