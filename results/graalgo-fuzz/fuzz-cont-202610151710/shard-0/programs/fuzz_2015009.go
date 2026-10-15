package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    v1 := []uint32{488, 2, 706, 244}
    var v2 uint64 = 362
    v3 := make(map[int]int)
    var v4 int32 = 379
    v5 := GenS1{F0: 144}
    v2 = 4294967295
    var v6 float64 = (float64((1e10 + ((3.14159 * (1.5 - 1e10)) + -2.25))) + 1.5)
    var v7 int32 = v4
    v8 := GenS1{F0: 24}
    for fi9 := 0; fi9 < 7; fi9++ {
        v7 = (^(v4))
        switch (fi9 >> ((255) & 63)) {
        case 95:
            for fi10 := 0; fi10 < 7; fi10++ {
                var v11 float32 = (-2.25 - (-2.25 * (0.0 - (3.14159 - 3.14159))))
                var v12 int32 = v7
                _ = fi10
                _ = v11
                _ = v12
            }
        default:
            v7 += v4
        }
        _ = fi9
    }
    if (true && false) {
        fmt.Println(v7)
        for fi13 := 0; fi13 < 14; fi13++ {
            var v14 uint32 = 4294967295
            var v15 uint64 = (^(v2))
            _ = fi13
            _ = v14
            _ = v15
        }
    }
    var v16 uint64 = v2
    switch 9223372036854775807 {
    case 0:
        for fi17 := 0; fi17 < 7; fi17++ {
            if len(v1) > 0 {
                v1[0] = 707
            }
            _ = fi17
        }
    }
    var v18 int = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
