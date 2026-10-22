package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 string }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[int]uint32)
    var v3 int = -2147483648
    var v4 uint64 = 42
    v4 -= (^(v4))
    v5 := GenS1{F0: 1, F1: 1, F2: "\n\t"}
    v4 -= (^(v4))
    v3 -= v3
    delete(v2, (-(v3)))
    var v6 float32 = 1e10
    if true {
        var v7 int8 = int8(v4)
        switch "\n\t" {
        case "hello":
            var v8 uint8 = 0
            _ = v8
        case "":
            v7 -= (v7 / ((-54) | 1))
        }
        var v9 uint8 = uint8(v4)
        _ = v7
        _ = v9
    }
    fmt.Println(v6, v3, v3)
    v4 = (^(v4))
    v4 = (^(v4))
    if ((v6 + -2.25) < (v6 / 1.5)) {
        v10 := GenS1{F0: 398, F1: 32767, F2: "hello"}
        var v11 int16 = -1
        _ = v10
        _ = v11
    } else {
        var v12 uint = 0
        var v13 int16 = -198
        var v14 int8 = 0
        _ = v12
        _ = v13
        _ = v14
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
