package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint }

func main() {
    var v1 uint = 255
    var v2 int8 = 100
    v3 := []int{9223372036854775807}
    var v4 uint16 = 42
    var v5 uint32 = 4294967295
    v6 := GenS1{F0: 2}
    if false {
        v7 := GenS1{F0: 1}
        _ = v7
    }
    var v8 bool = (true || true)
    v9 := GenS2{F0: uint(v4)}
    if len(v3) > 0 {
        v3[0] = 127
    }
    for fi10 := 0; fi10 < 20; fi10++ {
        v8 = false
        if (!true) {
            var v11 int = (-(fi10))
            fmt.Println(v8, v8)
            _ = v11
        }
        _ = fi10
    }
    for fi12 := 0; fi12 < 14; fi12++ {
        v2 = (v2 & (v2 / ((v2) | 1)))
        _ = fi12
    }
    for fi13 := 0; fi13 < 8; fi13++ {
        v1 = (v1 + v1)
        _ = fi13
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
