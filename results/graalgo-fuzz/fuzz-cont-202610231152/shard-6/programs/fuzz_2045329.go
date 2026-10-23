package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    var v1 uint32 = 2147483648
    var v2 uint = 1
    v3 := make(map[int]uint32)
    var v4 int64 = int64(v1)
    var v5 uint8 = 134
    v6 := make(map[int]int)
    var v7 float64 = (3.14159 + (-2.25 - (-2.25 / (0.1 / 3.14159))))
    var v8 int = 5
    if false {
        v4 -= (v4 + v4)
        var v9 int32 = -2147483648
        _ = v9
    }
    v4 = (v4 << (((v2 & v2)) & 63))
    v2 -= v2
    fmt.Println(v7)
    var v10 int16 = -1
    for fi11 := 0; fi11 < 3; fi11++ {
        if ((-(v4)) > v4) {
            for fi12 := 0; fi12 < 13; fi12++ {
                var v13 uint = v2
                _ = fi12
                _ = v13
            }
            v1 -= 2
            var v14 uint64 = 112
            _ = v14
        }
        _ = fi11
    }
    v15 := GenS1{F0: 1}
    v5 += v5
    v16 := v6[(v8 | v8)]
    v2 -= (^(v2))
    _ = len(v6)
    v3[128] = v1
    var v17 float32 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
