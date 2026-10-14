package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    var v1 int8 = 42
    v2 := make(map[uint32]string)
    var v3 uint = 403
    var v4 int = -2147483648
    var v5 uint8 = 1
    var v6 int32 = 2147483647
    v4 -= (v4 & (v4 | (v4 % ((v4) | 1))))
    var v7 float64 = 0.1
    var v8 uint32 = 4294967295
    var v9 float32 = float32(((3.14159 / -2.25) - v7))
    if (v9 != 1.5) {
        v10 := GenS1{F0: v3}
        _ = v10
    } else {
        if (760 <= -174) {
            v6 = (v6 << ((v3) & 63))
            var v11 uint64 = 1
            if ((1 <= -2147483648) && (!false)) {
                var v12 bool = false
                _ = v12
            } else {
                var v13 float64 = v7
                _ = v13
            }
            _ = v11
        } else {
            _ = len(v2)
            var v14 uint64 = 354
            for fi15 := 0; fi15 < 10; fi15++ {
                v16 := GenS1{F0: v3}
                _ = fi15
                _ = v16
            }
            _ = v14
        }
    }
    v17 := v2[(^(v8))]
    var v18 uint64 = 1
    var v19 int16 = -1
    v20 := []int{312, (v4 << ((65535) & 63)), 32767, v4}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
