package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    var v1 float32 = float32((-2.25 / 1e10))
    v2 := make(map[uint32]bool)
    var v3 uint = 65535
    if (32767 > 32767) {
        var v4 bool = false
        _ = v4
    }
    v5 := GenS1{F0: 0}
    fmt.Println(v3, v3)
    for fi6 := 0; fi6 < 9; fi6++ {
        v3 = v3
        _ = fi6
    }
    fmt.Println(v1, v3)
    switch 241 {
    case 4294967295:
        fmt.Println(v1)
    case 2147483647:
        if (0 < 1) {
            var v7 int16 = 1
            if (4294967295 < 2147483647) {
                var v8 int = 0
                _ = v8
            }
            _ = v7
        }
    case 370:
        v9 := []int{42}
        _ = v9
    default:
        if true {
            var v10 float64 = float64(float32((3.14159 - 1e10)))
            var v11 uint32 = 1
            _ = v10
            _ = v11
        }
    }
    v3 += v3
    v12 := v2[0]
    fmt.Println(v1, v3)
    v13 := GenS1{F0: -128}
    v12 = (true || (25 != 262))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v5
}
