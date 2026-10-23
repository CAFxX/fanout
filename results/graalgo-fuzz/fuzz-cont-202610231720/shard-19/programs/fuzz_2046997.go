package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    var v1 int16 = -593
    var v2 int = 32767
    v3 := make(map[string]bool)
    var v4 int8 = 42
    var v5 float64 = 0.1
    var v6 int64 = -775
    v7 := GenS1{F0: ((!((v2 - v2) <= (v2 | 0))) && (!(false && true)))}
    var v8 bool = ((43 > 255) || (!true))
    v9 := []int{32768, v2, (v2 >> ((4294967295) & 63))}
    v6 += (v6 % ((-9223372036854775808) | 1))
    v10 := GenS1{F0: v8}
    fmt.Println(v5)
    var v11 uint8 = 1
    v12 := v3["\n\t"]
    var v13 bool = true
    var v14 bool = v12
    v3["\n\t"] = false
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
