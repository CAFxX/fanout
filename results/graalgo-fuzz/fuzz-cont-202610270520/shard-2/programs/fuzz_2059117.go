package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    var v1 int8 = -128
    v2 := make(map[uint32]int)
    v3 := make(map[uint32]int)
    var v4 int32 = 1
    for fi5 := 0; fi5 < 6; fi5++ {
        v4 -= 42
        var v6 uint8 = 2
        var v7 int32 = v4
        _ = fi5
        _ = v6
        _ = v7
    }
    v8 := []int{0, -2147483648, 1, -840}
    switch 4294967295 {
    case -38:
        var v9 int = 9223372036854775807
        _ = v9
    }
    v1 += v1
    v4 -= (v4 >> ((1) & 63))
    var v10 float32 = (1.5 - (1e10 + 0.1))
    var v11 uint8 = 161
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    var v12 int8 = (v1 / ((-1) | 1))
    var v13 int64 = 42
    for fi14 := 0; fi14 < 19; fi14++ {
        v15 := GenS1{F0: 2}
        v13 -= v13
        if ((!(v12 == 1)) && true) {
            var v16 uint64 = 442
            var v17 uint32 = 65535
            _ = v16
            _ = v17
        } else {
            for fi18 := 0; fi18 < 12; fi18++ {
                if false {
                    switch (fi18 * (fi18 >> ((0) & 63))) {
                    case 127:
                        v1 = v1
                    case 949:
                        v19 := GenS1{F0: 65535}
                        _ = v19
                    default:
                        fmt.Println(v4, fi18, fi18)
                    }
                    v3[122] = 32767
                } else {
                    v20 := v2[300]
                    _ = v20
                }
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
                var v21 uint = 0
                _ = fi18
                _ = v21
            }
        }
        _ = fi14
        _ = v15
    }
    if len(v8) > 0 {
        v8[0] = 4294967295
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
