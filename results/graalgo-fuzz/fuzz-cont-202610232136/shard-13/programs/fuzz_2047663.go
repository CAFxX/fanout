package main

import (
    "fmt"
)


func main() {
    var v1 float32 = ((1.5 + 1e10) - 3.14159)
    v2 := make(map[string]uint32)
    var v3 uint = 887
    var v4 uint32 = 2
    var v5 int32 = 2147483647
    v1 = (((v1 * v1) / v1) / (1.5 - float32(float64(1e10))))
    var v6 int8 = 125
    var v7 uint64 = 578
    for fi8 := 0; fi8 < 9; fi8++ {
        var v9 int32 = (v5 % (((v5 % ((-2147483648) | 1))) | 1))
        var v10 uint32 = (v4 & uint32(v3))
        var v11 int64 = -214
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    var v12 int64 = 42
    fmt.Println(v3, v6, v7)
    for fi13 := 0; fi13 < 8; fi13++ {
        v14 := v2["hello"]
        v2[("foo bar" + ("\n\t" + "hello"))] = (^(v14))
        var v15 uint = (v3 & v3)
        _ = fi13
        _ = v14
        _ = v15
    }
    var v16 uint8 = 127
    v6 -= (v6 ^ v6)
    for fi17 := 0; fi17 < 4; fi17++ {
        v18 := []int{int(v3), fi17, (fi17 | fi17)}
        v12 = (v12 + (v12 | (v12 ^ v12)))
        var v19 bool = false
        _ = fi17
        _ = v18
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
