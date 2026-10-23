package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }

func main() {
    var v1 uint16 = 42
    v2 := make(map[string]string)
    v3 := []uint32{42}
    var v4 string = "\n\t"
    var v5 int64 = 9223372036854775807
    for fi6 := 0; fi6 < 14; fi6++ {
        var v7 int = fi6
        for fi8 := 0; fi8 < 1; fi8++ {
            v9 := v2["hello"]
            _ = fi8
            _ = v9
        }
        _ = fi6
        _ = v7
    }
    var v10 uint8 = 1
    v11 := GenS1{F0: 18446744073709551615, F1: 127}
    v5 = v5
    var v12 uint32 = 2
    if (-705 != -32768) {
        v4 = v4
        for fi13 := 0; fi13 < 13; fi13++ {
            v10 = (^(v10))
            _ = fi13
        }
        fmt.Println(v5, v10, v4)
    }
    v1 = 1
    v1 = 533
    var v14 int32 = 1
    for fi15 := 0; fi15 < 15; fi15++ {
        fmt.Println(v1)
        fmt.Println(fi15, v4, fi15)
        _ = fi15
    }
    fmt.Println(v1, v14)
    if (float32((1.5 + 0.0)) != 1.5) {
        v10 = (v10 << ((42) & 63))
        var v16 float32 = (float32(3.14159) / ((float32(1e10) / 1e10) * 3.14159))
        v16 += ((v16 + 0.1) + (float32(float64(0.1)) - v16))
        _ = v16
    } else {
        if true {
            _ = len(v2)
            var v17 int32 = v14
            var v18 uint32 = v12
            _ = v17
            _ = v18
        } else {
            var v19 int8 = int8(v12)
            fmt.Println(v5)
            _ = v19
        }
        var v20 uint32 = v12
        v21 := GenS1{F0: 804, F1: (v10 << ((463) & 63))}
        _ = v20
        _ = v21
    }
    fmt.Println(v10, v10)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
