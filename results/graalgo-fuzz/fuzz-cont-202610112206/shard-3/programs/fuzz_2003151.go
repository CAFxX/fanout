package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
func genFn0() int16 {
    var v5 string = ""
    v5 = ""
    _ = v5
    return 79
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint = 0
    var v3 float32 = 0.1
    var v4 int16 = -882
    var v6 uint = 1
    if (2147483647 <= 2147483647) {
        fmt.Println(v4)
    }
    if (-285 != -2147483648) {
        var v7 int16 = 42
        _ = v7
    } else {
        v6 -= (v6 - (v2 << (((v6 >> ((18446744073709551615) & 63))) & 63)))
        fmt.Println(v4, v2, v3)
    }
    for fi8 := 0; fi8 < 19; fi8++ {
        v3 -= 0.0
        v3 = v3
        _ = fi8
    }
    var v9 uint = v6
    v10 := GenS1{F0: 42, F1: 42}
    v6 += v9
    _ = len(v1)
    v3 = (v3 + 3.14159)
    var v11 bool = true
    var v12 float64 = 1e10
    for fi13 := 0; fi13 < 4; fi13++ {
        var v14 int8 = -1
        v6 = 584
        var v15 int16 = v4
        _ = fi13
        _ = v14
        _ = v15
    }
    v16 := GenS1{F0: -741, F1: 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v9
}
