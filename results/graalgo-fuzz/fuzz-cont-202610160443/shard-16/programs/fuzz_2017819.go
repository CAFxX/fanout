package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32; F2 uint8 }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[uint32]uint32)
    v3 := make(map[uint32]uint32)
    var v4 int16 = -1
    var v5 uint16 = 2
    fmt.Println(v5, v4)
    v4 = 237
    var v6 uint64 = 1
    v4 = 32767
    delete(v3, 0)
    v7 := GenS1{F0: true, F1: 2147483647, F2: 128}
    switch "\n\t" {
    case "\n\t":
        v4 -= v4
    case "":
        if (-2.25 <= (1e10 / (1.5 / (0.0 + (float32(float64(3.14159)) - -2.25))))) {
            var v8 int16 = (^(v4))
            v9 := []int{32767}
            _ = v8
            _ = v9
        } else {
            v6 = (v6 % ((v6) | 1))
            v6 = (v6 - v6)
        }
    case "a":
        fmt.Println(v6)
    }
    v6 += (v6 << ((65535) & 63))
    var v10 int8 = 100
    v10 += 42
    if false {
        v6 -= 4294967295
        fmt.Println(v5)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
