package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    var v1 float32 = 3.14159
    var v2 int64 = 9223372036854775807
    var v3 int64 = v2
    v4 := []int{1}
    fmt.Println(v1)
    var v5 int16 = -586
    var v6 bool = (-1 != 1)
    for fi7 := 0; fi7 < 11; fi7++ {
        var v8 int16 = v5
        var v9 uint64 = 1
        for fi10 := 0; fi10 < 1; fi10++ {
            v11 := GenS1{F0: 2147483648}
            _ = fi10
            _ = v11
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    var v12 int = 32768
    if (3.14159 <= float32((1.5 * float64((float32(3.14159) + v1))))) {
        if len(v4) > 0 {
            v4[0] = v12
        }
    } else {
        _ = len(v4)
    }
    if len(v4) > 0 {
        v4[0] = int(v2)
    }
    fmt.Println(v6)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
