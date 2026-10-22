package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
func genFn0(p0 uint8) int8 {
    var v6 uint8 = (^(p0))
    for fi7 := 0; fi7 < 8; fi7++ {
        for fi8 := 0; fi8 < 18; fi8++ {
            v9 := GenS1{F0: 255, F1: true}
            if (("foo bar" >= "\n\t") || (1 != 2147483648)) {
                var v10 uint8 = 227
                _ = v10
            }
            p0 += p0
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    _ = p0
    _ = v6
    return 100
}
func genFn1() int32 {
    v11 := []int{-554, -2147483648}
    var v12 int64 = -297
    _ = v11
    _ = v12
    return 0
}

func main() {
    var v1 float64 = ((1e10 - float64((((0.0 * 0.1) * 3.14159) - 1e10))) + float64(float32(((3.14159 / (0.1 * float64((float32(1e10) - 1.5)))) - 1e10))))
    var v2 uint32 = 848
    v3 := []int{-9223372036854775808, 2147483647, 2147483647}
    var v4 uint8 = 42
    v5 := []string{"hello", "hello", ""}
    for fi13 := 0; fi13 < 6; fi13++ {
        fmt.Println(fi13)
        v4 = v4
        _ = fi13
    }
    var v14 int = 947
    var v15 uint32 = 65535
    v16 := GenS1{F0: 0, F1: (false || (!false))}
    var v17 uint8 = v4
    var v18 int = 1
    var v19 uint64 = 4294967295
    if len(v5) > 0 {
        v5[0] = ("foo bar" + "\n\t")
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
