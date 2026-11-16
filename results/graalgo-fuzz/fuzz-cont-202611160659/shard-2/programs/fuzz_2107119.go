package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint16; F1 uint16 }

func main() {
    v1 := []uint32{2147483647, 1, 2147483648, 4294967295}
    var v2 float32 = (0.0 * -2.25)
    var v3 int = 0
    var v4 int = 32768
    var v5 uint64 = 2
    if (1 != 42) {
        var v6 uint32 = 0
        for fi7 := 0; fi7 < 1; fi7++ {
            v5 = v5
            var v8 bool = (2147483647 < 1)
            var v9 int64 = -104
            _ = fi7
            _ = v8
            _ = v9
        }
        var v10 uint = 0
        _ = v10
        _ = v6
    } else {
        if true {
            v11 := GenS1{F0: true}
            v2 -= 0.1
            if (true || (1 != 1)) {
                if len(v1) > 0 {
                    v1[0] = 15
                }
                _ = len(v1)
            }
            _ = v11
        } else {
            var v12 int32 = -2147483648
            _ = v12
        }
    }
    fmt.Println(v3, v4)
    for fi13 := 0; fi13 < 17; fi13++ {
        var v14 uint32 = 4294967295
        var v15 uint32 = 4294967295
        _ = fi13
        _ = v14
        _ = v15
    }
    v5 += (^(v5))
    var v16 int = (v3 - 2147483647)
    var v17 int8 = 122
    _ = len(v1)
    v18 := GenS2{F0: 518, F1: 42}
    var v19 int = (v16 - 2147483647)
    fmt.Println(v16, v5, v17)
    fmt.Println(v2, v17, v2)
    v16 += 216
    var v20 uint8 = 42
    v19 -= v16
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
