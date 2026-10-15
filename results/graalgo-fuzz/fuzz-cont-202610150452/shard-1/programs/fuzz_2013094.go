package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32; F2 uint64 }
func genFn0(p0 int32, p1 int32, p2 int) uint8 {
    if ((1 == 42) && false) {
        for fi6 := 0; fi6 < 3; fi6++ {
            var v7 uint16 = 0
            v7 += (v7 << ((0) & 63))
            _ = fi6
            _ = v7
        }
    } else {
        v8 := []int{}
        _ = v8
    }
    p0 -= (p0 | p0)
    p2 += p2
    _ = p0
    _ = p1
    _ = p2
    return 42
}

func main() {
    var v1 uint16 = 65535
    var v2 bool = (0 == -32768)
    var v3 int8 = 100
    var v4 int = 42
    var v5 uint16 = v1
    var v9 int = v4
    var v10 float32 = 1.5
    v5 = v1
    var v11 int = (v9 + -650)
    v2 = v2
    v1 += (^(v5))
    v10 = (1e10 - v10)
    v12 := GenS1{F0: (42 >= 121), F1: -2147483648, F2: 790}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
