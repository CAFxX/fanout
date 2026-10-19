package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    var v1 int8 = -117
    var v2 int16 = -32768
    v3 := []int{}
    if len(v3) > 0 {
        v3[0] = 654
    }
    fmt.Println(v1, v2)
    var v4 uint8 = 177
    if (1 == -784) {
        v2 += v2
        _ = len(v3)
    } else {
        v2 += (^(v2))
        v5 := GenS1{F0: 42}
        for fi6 := 0; fi6 < 5; fi6++ {
            var v7 int = fi6
            _ = fi6
            _ = v7
        }
        _ = v5
    }
    for fi8 := 0; fi8 < 11; fi8++ {
        for fi9 := 0; fi9 < 20; fi9++ {
            var v10 float64 = 3.14159
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    var v11 int16 = 1000
    v4 += v4
    if false {
        var v12 uint16 = 42
        _ = v12
    } else {
        fmt.Println(v11, v4, v2)
        var v13 bool = true
        _ = v13
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
}
