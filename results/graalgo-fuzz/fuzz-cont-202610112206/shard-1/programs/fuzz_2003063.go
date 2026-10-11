package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8 }

func main() {
    var v1 float64 = (0.1 / 1e10)
    var v2 int16 = 321
    v3 := []string{"\n\t", "\n\t", "\n\t"}
    var v4 string = ""
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v5 uint64 = 0
    var v6 int = -9223372036854775808
    v7 := GenS1{F0: 100}
    v2 += 1000
    v8 := GenS2{F0: -28}
    v2 = 1000
    var v9 uint8 = 0
    var v10 int8 = 1
    if true {
        v3 = append(v3, v4)
        v4 = v4
        v1 -= v1
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
