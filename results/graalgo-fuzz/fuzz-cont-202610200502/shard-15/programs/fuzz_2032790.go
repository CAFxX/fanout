package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    var v1 int = 423
    var v2 int16 = -32768
    var v3 int64 = -9223372036854775808
    var v4 int = (v1 + v1)
    var v5 int64 = v3
    if true {
        var v6 int64 = (v3 >> ((42) & 63))
        v7 := GenS1{F0: v2}
        _ = v6
        _ = v7
    }
    if (4294967295 < 1) {
        var v8 int16 = (v2 << ((65535) & 63))
        v2 = (-(v8))
        if true {
            for fi9 := 0; fi9 < 12; fi9++ {
                v10 := GenS1{F0: (^(v2))}
                if true {
                    var v11 uint8 = 187
                    _ = v11
                }
                _ = fi9
                _ = v10
            }
        }
        _ = v8
    }
    var v12 float64 = (float64(3.14159) - 1.5)
    v13 := []int{(v4 % ((v4) | 1))}
    var v14 uint8 = 124
    v2 = (v2 & 686)
    if false {
        var v15 int16 = v2
        v12 -= (-2.25 - 1.5)
        v3 -= 1
        _ = v15
    }
    v16 := GenS1{F0: v2}
    if len(v13) > 0 {
        v13[0] = (v4 & (v4 ^ v4))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
