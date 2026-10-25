package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    v1 := make(map[string]string)
    var v2 uint16 = 255
    v3 := make(map[int]bool)
    var v4 uint16 = 1
    v5 := make(map[uint32]string)
    v6 := make(map[uint32]string)
    fmt.Println(v2, v4)
    for fi7 := 0; fi7 < 18; fi7++ {
        v4 = 165
        v5[1] = ("" + (("a" + "foo bar") + "\n\t"))
        _ = fi7
    }
    v4 = (v4 / ((v2) | 1))
    v8 := GenS1{F0: 4294967295}
    fmt.Println(v2)
    var v9 int32 = 42
    var v10 float64 = -2.25
    var v11 int16 = 42
    var v12 float32 = (1e10 / (3.14159 + ((1e10 - float32(1e10)) - 1e10)))
    v2 -= 550
    var v13 float32 = float32(((3.14159 / (((v10 * (3.14159 + 1e10)) / v10) * v10)) - 1e10))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
