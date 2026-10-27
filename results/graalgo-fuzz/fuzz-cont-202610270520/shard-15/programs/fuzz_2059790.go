package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
type GenS2 struct { F0 uint32 }

func main() {
    var v1 float32 = (0.1 * 0.1)
    var v2 uint64 = 238
    var v3 bool = true
    v4 := []string{}
    var v5 float64 = (0.1 / 1.5)
    var v6 int = 679
    var v7 string = (("" + ("hello" + "hello")) + ("a" + (("a" + (("a" + "a") + ("\n\t" + "\n\t"))) + "")))
    v4 = append(v4, v7)
    v2 += v2
    switch v6 {
    case 952:
        v8 := GenS1{F0: 133, F1: ((1 > 255) || v3)}
        _ = v8
    }
    v9 := GenS2{F0: 1}
    v1 -= 3.14159
    var v10 string = ""
    v5 -= v5
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
