package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
func genFn0(p0 uint64, p1 bool, p2 uint32) uint32 {
    v5 := []int{1, 42, 124, -9223372036854775808}
    fmt.Println(p2, p2)
    for fi6 := 0; fi6 < 10; fi6++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        p1 = p1
        _ = fi6
    }
    _ = len(v5)
    var v7 float32 = float32(0.1)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v7
    return (p2 / ((42) | 1))
}

func main() {
    var v1 int = 527
    v2 := make(map[string]uint32)
    var v3 float32 = 1e10
    var v4 int8 = -1
    var v8 uint8 = 42
    v9 := GenS1{F0: -2147483648, F1: 127}
    v3 -= 1.5
    var v10 int16 = 1
    v1 -= (v1 >> ((4294967295) & 63))
    fmt.Println(v1, v10, v3)
    for fi11 := 0; fi11 < 11; fi11++ {
        fmt.Println(v8, v1)
        var v12 int64 = 0
        if (true || ((!(18446744073709551615 > 470)) || false)) {
            v4 = (v4 ^ 1)
            v3 += v3
            var v13 uint64 = 18446744073709551615
            _ = v13
        }
        _ = fi11
        _ = v12
    }
    for fi14 := 0; fi14 < 18; fi14++ {
        var v15 uint = 18446744073709551615
        _ = fi14
        _ = v15
    }
    v3 = -2.25
    v4 = 100
    var v16 int = (v1 ^ (v1 % ((v1) | 1)))
    v8 = 0
    var v17 int32 = 0
    if false {
        v2["\n\t"] = 516
        fmt.Println(v16, v1, v10)
        var v18 int = (v1 ^ v1)
        _ = v18
    }
    v16 = v16
    var v19 int16 = v10
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
