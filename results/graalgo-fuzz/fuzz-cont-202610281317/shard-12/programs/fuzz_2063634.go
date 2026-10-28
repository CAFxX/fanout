package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }

func main() {
    var v1 bool = (!((839 >= 0) || true))
    var v2 int32 = 2147483647
    v3 := []int{4294967295, 1, -2147483648, int(v2)}
    var v4 uint = 2
    v2 = v2
    if false {
        if (!true) {
            fmt.Println(v1, v2, v2)
            v2 -= 1
            if true {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                v5 := GenS1{F0: 0, F1: v4}
                for fi6 := 0; fi6 < 11; fi6++ {
                    v7 := GenS1{F0: 2, F1: (v4 & (v4 << ((0) & 63)))}
                    _ = fi6
                    _ = v7
                }
                _ = v5
            }
        }
    }
    var v8 int8 = 73
    var v9 int = -735
    var v10 float32 = ((0.0 * float32((0.0 + float64(3.14159)))) - float32(3.14159))
    var v11 int16 = 0
    fmt.Println(v4, v9, v2)
    fmt.Println(v2, v9)
    v10 -= 1e10
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
