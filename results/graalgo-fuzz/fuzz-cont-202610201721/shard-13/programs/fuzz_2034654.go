package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0(p0 int8, p1 int8, p2 string) string {
    v5 := []int{}
    if true {
        var v6 int8 = p1
        var v7 bool = true
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = v6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return p2
}

func main() {
    var v1 uint = 18446744073709551615
    var v2 uint8 = 218
    var v3 int = -467
    var v4 uint8 = (^(v2))
    v2 = 133
    var v8 uint32 = 1
    v4 -= (v4 / ((84) | 1))
    v9 := GenS1{F0: (v2 << ((uint(v8)) & 63))}
    v1 -= 0
    var v10 uint32 = 0
    for fi11 := 0; fi11 < 1; fi11++ {
        v12 := GenS1{F0: (v2 | (v4 * (v2 & v2)))}
        _ = fi11
        _ = v12
    }
    var v13 int32 = -259
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
