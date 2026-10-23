package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 bool; F1 string; F2 uint64 }

func main() {
    v1 := make(map[string]bool)
    var v2 int16 = 1000
    v3 := make(map[string]bool)
    var v4 int8 = -128
    v5 := make(map[string]bool)
    var v6 float64 = 0.0
    var v7 uint8 = 1
    for fi8 := 0; fi8 < 16; fi8++ {
        v2 = v2
        v9 := GenS1{F0: 18446744073709551615, F1: 689}
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 10; fi10++ {
        v4 -= (v4 - int8(v2))
        _ = fi10
    }
    var v11 int64 = 495
    v3["\n\t"] = true
    v11 = 1
    v12 := GenS2{F0: (-1 >= 42), F1: (("a" + "") + ""), F2: 42}
    v13 := GenS2{F0: true, F1: "foo bar", F2: 201}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
