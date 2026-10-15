package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    var v1 bool = true
    v2 := []int{-9223372036854775808}
    var v3 uint64 = 4
    for fi4 := 0; fi4 < 3; fi4++ {
        var v5 int = fi4
        _ = fi4
        _ = v5
    }
    v6 := GenS1{F0: v3}
    switch 32768 {
    case -2147483648:
        v1 = false
    }
    if len(v2) > 0 {
        v2[0] = -1
    }
    var v7 float32 = -2.25
    var v8 uint8 = 174
    v8 = 6
    v9 := GenS1{F0: (v3 * (v3 - 646))}
    fmt.Println(v8, v3, v7)
    for fi10 := 0; fi10 < 13; fi10++ {
        for fi11 := 0; fi11 < 4; fi11++ {
            var v12 float64 = ((0.1 * 0.1) + (0.1 * 1e10))
            _ = fi11
            _ = v12
        }
        _ = fi10
    }
    if true {
        fmt.Println(v1, v7)
        v13 := GenS1{F0: v3}
        v1 = (((((v7 + 3.14159) - (0.0 * float32((1e10 * 0.1)))) + (((1.5 - float32(0.1)) + 0.0) + 1e10)) < 1.5) && false)
        _ = v13
    }
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
