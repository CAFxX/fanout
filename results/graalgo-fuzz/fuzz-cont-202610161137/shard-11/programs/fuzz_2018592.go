package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 uint32 }
type GenS2 struct { F0 uint64; F1 uint }

func main() {
    var v1 uint = 1
    v2 := make(map[string]uint32)
    v3 := make(map[uint32]uint32)
    var v4 int = 42
    v4 -= 161
    var v5 string = "a"
    v3[65535] = 2
    switch v4 {
    case 128:
        switch ("foo bar" + v5) {
        case "hello":
            fmt.Println(v5, v1)
        case "a":
            if (!false) {
                fmt.Println(v5, v4)
                fmt.Println(v5, v5)
            }
        }
    }
    var v6 float32 = 0.0
    v7 := GenS1{F0: 42, F1: (v4 % ((v4) | 1)), F2: 2147483648}
    v4 += 127
    v8 := GenS2{F0: 0, F1: v1}
    if false {
        v9 := GenS1{F0: 2, F1: (v4 | (v4 << ((42) & 63))), F2: 261}
        v10 := GenS1{F0: 42, F1: 3, F2: 0}
        _ = v10
        _ = v9
    }
    v1 -= v1
    var v11 bool = true
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
