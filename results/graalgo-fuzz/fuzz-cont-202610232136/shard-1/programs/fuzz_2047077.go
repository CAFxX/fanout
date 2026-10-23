package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    var v1 bool = false
    var v2 uint8 = 42
    var v3 int32 = 974
    v2 -= (^(v2))
    v4 := []int{1, 4294967295, 789, -1}
    var v5 int = 4294967295
    for fi6 := 0; fi6 < 9; fi6++ {
        if (((v2 + v2) == 100) && (v1 && (!(4294967295 == 410)))) {
            v7 := GenS1{F0: 9223372036854775807}
            fmt.Println(v5, fi6)
            _ = v7
        }
        _ = fi6
    }
    v5 = -915
    if v1 {
        for fi8 := 0; fi8 < 3; fi8++ {
            v5 = fi8
            for fi9 := 0; fi9 < 9; fi9++ {
                fmt.Println(v5, fi9, v2)
                _ = fi9
            }
            var v10 float32 = 0.1
            _ = fi8
            _ = v10
        }
    }
    var v11 int16 = 1000
    v12 := GenS1{F0: -9223372036854775808}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
