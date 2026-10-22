package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }

func main() {
    v1 := make(map[uint32]string)
    var v2 int16 = -32768
    var v3 int32 = 0
    v4 := make(map[string]int)
    var v5 string = "a"
    var v6 uint16 = 159
    var v7 int = 42
    v7 = -2147483648
    v3 = (v3 | (-(v3)))
    v8 := GenS1{F0: 100, F1: 1}
    v5 = "\n\t"
    v7 = int(v2)
    switch 4294967295 {
    case -546:
        v9 := GenS1{F0: 113, F1: 0}
        _ = v9
    case 42:
        if false {
            v10 := v4["a"]
            var v11 bool = true
            _ = v10
            _ = v11
        }
    case 551:
        switch v7 {
        case 128:
            var v12 uint8 = 1
            _ = v12
        }
    }
    v6 = v6
    v6 += v6
    var v13 float32 = 1.5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
