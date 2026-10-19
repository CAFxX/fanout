package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }

func main() {
    v1 := []uint32{2, 2147483648, 65535}
    var v2 float32 = -2.25
    var v3 uint = 1
    var v4 int32 = -1
    v5 := make(map[uint32]int)
    v6 := make(map[string]uint32)
    var v7 int32 = v4
    fmt.Println(v3, v7)
    var v8 string = "foo bar"
    fmt.Println(v2, v3, v7)
    fmt.Println(v7, v2)
    if (!((784 != 88) && true)) {
        if true {
            v3 -= (v3 ^ v3)
            v3 -= (v3 + 352)
            v9 := GenS1{F0: -2147483648, F1: 1000}
            _ = v9
        } else {
            switch 32768 {
            case 1:
                var v10 uint16 = 156
                _ = v10
            case 32768:
                v8 = "a"
            }
        }
        v3 += (v3 | 1)
        delete(v6, v8)
    }
    v2 += (0.1 * (v2 * v2))
    if len(v1) > 0 {
        v1[0] = 0
    }
    v11 := v5[65535]
    v7 = -2147483648
    var v12 int = (v11 - (v11 % ((v11) | 1)))
    v4 = v4
    for fi13 := 0; fi13 < 2; fi13++ {
        if false {
            v2 = v2
            if len(v1) > 0 {
                v1[0] = 42
            }
            v5[65535] = 1
        } else {
            var v14 float32 = ((v2 / (v2 - 0.0)) + v2)
            _ = v14
        }
        _ = fi13
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
