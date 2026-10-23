package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
func genFn0(p0 uint16, p1 uint8, p2 uint) bool {
    p2 = p2
    p2 = 18446744073709551615
    _ = p0
    _ = p1
    _ = p2
    return true
}

func main() {
    var v1 uint64 = 0
    var v2 int16 = 957
    var v3 int64 = 698
    var v4 float64 = 1.5
    v4 += v4
    v3 = (v3 << ((687) & 63))
    fmt.Println(v4, v2)
    var v5 uint = 18446744073709551615
    v5 = (v5 | (^(v5)))
    fmt.Println(v3)
    var v6 bool = true
    v7 := GenS1{F0: 289, F1: 2}
    for fi8 := 0; fi8 < 10; fi8++ {
        var v9 bool = v6
        for fi10 := 0; fi10 < 17; fi10++ {
            v11 := GenS1{F0: v1, F1: 0}
            v1 -= 990
            _ = fi10
            _ = v11
        }
        for fi12 := 0; fi12 < 5; fi12++ {
            v9 = (v9 || false)
            _ = fi12
        }
        _ = fi8
        _ = v9
    }
    var v13 int32 = 42
    var v14 uint16 = 350
    var v15 uint16 = (v14 % (((v14 + (^(v14)))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
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
}
