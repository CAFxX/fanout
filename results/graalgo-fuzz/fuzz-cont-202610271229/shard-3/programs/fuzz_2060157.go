package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16 }
func genFn0(p0 uint32, p1 uint8, p2 uint16) uint8 {
    v5 := GenS1{F0: (p2 | p2)}
    p2 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 128
}

func main() {
    var v1 uint = 65535
    v2 := make(map[string]bool)
    var v3 int8 = 42
    var v4 float64 = 0.1
    var v6 int = 2147483647
    v7 := GenS2{F0: -1}
    v8 := GenS1{F0: 32767}
    v2["\n\t"] = true
    v4 -= float64(0.0)
    v4 += v4
    fmt.Println(v1, v4, v6)
    fmt.Println(v6, v3, v1)
    v9 := v2[""]
    var v10 int16 = -812
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
