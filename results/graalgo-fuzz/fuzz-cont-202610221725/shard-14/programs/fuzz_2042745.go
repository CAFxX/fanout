package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    var v1 uint16 = 785
    var v2 int = 128
    v3 := make(map[int]int)
    var v4 string = "foo bar"
    var v5 uint = 18446744073709551615
    var v6 uint64 = uint64(v1)
    v7 := GenS1{F0: -9223372036854775808}
    v8 := GenS1{F0: 9223372036854775807}
    var v9 uint = (v5 ^ v5)
    var v10 float64 = (3.14159 * (1e10 / 1.5))
    v11 := v3[(v2 + (v2 & int(v5)))]
    var v12 float64 = -2.25
    fmt.Println(v2)
    fmt.Println(v2)
    _ = len(v3)
    var v13 int8 = 26
    v14 := v3[(v11 + (v2 >> (((v5 | uint(v1))) & 63)))]
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
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
