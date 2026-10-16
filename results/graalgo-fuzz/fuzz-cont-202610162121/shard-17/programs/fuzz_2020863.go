package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    var v1 string = "hello"
    var v2 uint32 = 0
    var v3 int64 = -206
    var v4 int = -9223372036854775808
    var v5 float64 = ((3.14159 - 1e10) + ((3.14159 * -2.25) + (1.5 - 3.14159)))
    switch v4 {
    case -158:
        v6 := []int{32767}
        _ = v6
    }
    for fi7 := 0; fi7 < 16; fi7++ {
        v1 = v1
        _ = fi7
    }
    v4 = (^(v4))
    v8 := []int{v4, 9223372036854775807, v4, (v4 * v4)}
    var v9 int8 = 0
    if true {
        for fi10 := 0; fi10 < 14; fi10++ {
            var v11 float64 = v5
            var v12 float64 = ((float64(-2.25) * -2.25) + ((1e10 * (v5 / v5)) / v11))
            v1 = v1
            _ = fi10
            _ = v11
            _ = v12
        }
        fmt.Println(v1, v9)
        var v13 uint16 = 42
        _ = v13
    }
    var v14 int64 = 42
    v4 = v4
    if (0 < 813) {
        var v15 uint32 = v2
        v9 += v9
        fmt.Println(v14, v4)
        _ = v15
    }
    v9 -= v9
    v5 = v5
    v16 := GenS1{F0: v14}
    var v17 uint32 = (v2 & 2147483648)
    if len(v8) > 0 {
        v8[0] = v4
    }
    var v18 int32 = -2147483648
    fmt.Println(v14, v9, v14)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
