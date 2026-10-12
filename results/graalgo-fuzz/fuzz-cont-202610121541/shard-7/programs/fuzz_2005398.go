package main

import (
    "fmt"
)


func main() {
    var v1 uint8 = 66
    var v2 uint32 = 2147483647
    var v3 bool = (-128 <= 0)
    var v4 uint32 = 2147483647
    for fi5 := 0; fi5 < 6; fi5++ {
        var v6 int64 = int64(fi5)
        _ = fi5
        _ = v6
    }
    if (0.1 > 1.5) {
        v7 := []int{}
        fmt.Println(v3, v1, v1)
        for fi8 := 0; fi8 < 4; fi8++ {
            v4 = v2
            if len(v7) > 0 {
                v7[0] = (fi8 >> ((18446744073709551615) & 63))
            }
            _ = fi8
        }
        _ = v7
    } else {
        for fi9 := 0; fi9 < 20; fi9++ {
            var v10 int32 = -1
            var v11 uint64 = uint64(v4)
            var v12 int16 = -1
            _ = fi9
            _ = v10
            _ = v11
            _ = v12
        }
        var v13 float64 = 0.0
        _ = v13
    }
    var v14 float32 = (-2.25 / (3.14159 / ((1e10 - 0.1) + 3.14159)))
    fmt.Println(v4, v4)
    var v15 uint64 = 4294967295
    var v16 int16 = 32767
    var v17 float64 = -2.25
    v4 = (v2 + v2)
    v4 = 2147483648
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
