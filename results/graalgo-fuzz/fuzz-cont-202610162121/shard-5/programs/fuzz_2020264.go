package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0() int16 {
    var v7 bool = (2147483647 != 2147483647)
    v8 := GenS1{F0: 32767}
    _ = v7
    _ = v8
    return 42
}

func main() {
    var v1 bool = true
    var v2 uint8 = 237
    var v3 float64 = 3.14159
    v4 := make(map[int]string)
    var v5 uint8 = 127
    v6 := make(map[int]uint32)
    if (!((1 < 42) || false)) {
        v5 = (^(v5))
    } else {
        var v9 int64 = 145
        _ = v9
    }
    v10 := GenS1{F0: int16(v5)}
    v11 := v4[127]
    var v12 bool = v1
    for fi13 := 0; fi13 < 6; fi13++ {
        v1 = (v11 == v11)
        v2 = v5
        _ = fi13
    }
    var v14 uint64 = 1
    for fi15 := 0; fi15 < 9; fi15++ {
        fmt.Println(v5, v11)
        if (2 > 42) {
            fmt.Println(v5, v11, v11)
        }
        if v1 {
            fmt.Println(v14, v12)
        }
        _ = fi15
    }
    v2 += v5
    var v16 uint64 = (v14 | v14)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
