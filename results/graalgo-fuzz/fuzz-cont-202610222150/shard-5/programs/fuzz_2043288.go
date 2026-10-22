package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    var v1 int = -9223372036854775808
    var v2 uint8 = 127
    var v3 int8 = 1
    var v4 uint8 = v2
    v3 = v3
    v4 -= (v4 ^ 1)
    v4 -= (v2 ^ (^(v2)))
    var v5 int64 = 1
    for fi6 := 0; fi6 < 12; fi6++ {
        for fi7 := 0; fi7 < 6; fi7++ {
            if true {
                fmt.Println(v5)
                var v8 uint16 = 42
                var v9 uint8 = 251
                v10 := []int{(fi6 ^ (^(v1))), (fi7 % ((127) | 1)), 594}
                _ = v10
                _ = v8
                _ = v9
            } else {
                v11 := []int{fi6, int(v4), (^(fi7)), fi7}
                _ = len(v11)
                _ = v11
            }
            v12 := []int{(v1 >> ((488) & 63)), (v1 ^ (fi7 >> ((1) & 63)))}
            fmt.Println(fi7, v1, v2)
            _ = fi7
            _ = v12
        }
        _ = fi6
    }
    v13 := GenS1{F0: 18446744073709551615}
    fmt.Println(v1, v4, v1)
    var v14 float32 = 1.5
    if false {
        var v15 bool = (!true)
        _ = v15
    }
    v16 := []int{v1, v1, 0, (v1 | -790)}
    if true {
        v3 -= v3
    }
    v16 = append(v16, (v1 ^ (v1 * v1)))
    v14 = float32(0.1)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
