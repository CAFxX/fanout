package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0() uint8 {
    var v7 int = 42
    if (-1 == 176) {
        var v8 uint16 = 42
        v9 := GenS1{F0: 258}
        v10 := GenS1{F0: 42}
        _ = v10
        _ = v8
        _ = v9
    }
    var v11 bool = (1.5 > 1.5)
    _ = v11
    _ = v7
    return 22
}

func main() {
    var v1 bool = true
    var v2 float32 = (1e10 - ((-2.25 / (1.5 / (3.14159 * 0.1))) * 1.5))
    var v3 int = -383
    var v4 int64 = 1
    var v5 float64 = 1.5
    v6 := make(map[string]uint32)
    if ("foo bar" >= "\n\t") {
        v12 := []int{v3, 128, -9223372036854775808, (v3 - (^(v3)))}
        _ = v12
    }
    for fi13 := 0; fi13 < 7; fi13++ {
        var v14 float64 = 0.1
        _ = fi13
        _ = v14
    }
    v2 = v2
    if ((65535 <= 42) && true) {
        v15 := []int{-9223372036854775808, v3, (v3 >> ((644) & 63))}
        _ = v15
    }
    var v16 uint = 2
    var v17 float32 = v2
    v18 := GenS1{F0: 2147483647}
    var v19 int32 = 250
    if v1 {
        v20 := GenS1{F0: 2147483648}
        var v21 float32 = v2
        for fi22 := 0; fi22 < 19; fi22++ {
            v23 := v6["foo bar"]
            _ = fi22
            _ = v23
        }
        _ = v20
        _ = v21
    }
    var v24 int32 = -666
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
