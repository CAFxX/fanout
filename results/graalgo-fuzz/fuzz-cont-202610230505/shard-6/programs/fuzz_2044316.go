package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 int16 }

func main() {
    var v1 float64 = float64((1.5 / 1.0))
    var v2 bool = (!true)
    v3 := []string{"hello", (("foo bar" + "") + ("foo bar" + ("hello" + ("foo bar" + "\n\t"))))}
    var v4 float64 = float64(3.14159)
    v1 += ((0.1 - (float64(-2.25) + (0.0 - 0.1))) - v4)
    v4 += v4
    v2 = v2
    v4 += (((v4 - v4) / 3.14159) / float64((float32(v1) / 0.1)))
    fmt.Println(v4)
    v5 := GenS1{F0: 65535, F1: 128, F2: 1}
    v2 = false
    fmt.Println(v1, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
