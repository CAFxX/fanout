package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 int16 }
func genFn0() int {
    var v7 uint = 255
    var v8 uint64 = 42
    _ = v7
    _ = v8
    return 4294967295
}

func main() {
    var v1 int32 = 0
    v2 := make(map[int]int)
    var v3 uint8 = 255
    v4 := []uint32{144}
    var v5 float32 = 0.1
    var v6 uint32 = 65535
    var v9 int = -9223372036854775808
    var v10 uint16 = 511
    delete(v2, (v9 << ((166) & 63)))
    v3 = v3
    v3 += (v3 << ((42) & 63))
    var v11 bool = (62 == -1)
    v11 = true
    switch (v9 * (v9 / ((v9) | 1))) {
    case -2147483648:
        for fi12 := 0; fi12 < 3; fi12++ {
            for fi13 := 0; fi13 < 3; fi13++ {
                v14 := GenS1{F0: 255, F1: int16(v6), F2: -1}
                delete(v2, fi12)
                switch v9 {
                case 4294967295:
                    v9 -= 1
                }
                _ = fi13
                _ = v14
            }
            v1 = v1
            v15 := GenS1{F0: v10, F1: 1000, F2: 1}
            _ = fi12
            _ = v15
        }
    case -1:
        var v16 int8 = -64
        _ = v16
    }
    _ = len(v4)
    var v17 int = 9223372036854775807
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
