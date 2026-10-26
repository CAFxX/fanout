package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 int }

func main() {
    v1 := make(map[int]uint32)
    var v2 int = -2147483648
    v3 := make(map[int]uint32)
    var v4 int16 = -32768
    v5 := GenS1{F0: (0.1 < (1.5 + ((0.1 * 3.14159) + 0.0))), F1: -1}
    var v6 int64 = 1
    fmt.Println(v2)
    v6 += (v6 | 1)
    var v7 uint = 65535
    if true {
        v8 := []int{(v2 >> ((0) & 63)), 42}
        var v9 int32 = -1
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: true, F1: v6}
    var v11 int8 = -126
    var v12 int32 = 0
    switch (v2 & v2) {
    case -749:
        v13 := GenS2{F0: (-(v2))}
        _ = v13
    case 32767:
        if (!true) {
            for fi14 := 0; fi14 < 18; fi14++ {
                var v15 int32 = (v12 + v12)
                _ = fi14
                _ = v15
            }
            if (true && ((!true) || (!false))) {
                var v16 uint = (^(v7))
                _ = v16
            }
            if true {
                v6 += (v6 / ((v6) | 1))
                v17 := GenS1{F0: true, F1: v6}
                v4 = v4
                _ = v17
            }
        }
    case 127:
        v18 := GenS2{F0: (v2 + 127)}
        _ = v18
    default:
        var v19 string = ("a" + "a")
        _ = v19
    }
    v20 := GenS1{F0: (v6 > int64(v7)), F1: (v6 ^ int64(v4))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
