package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    v1 := make(map[int]string)
    var v2 int64 = 854
    v3 := make(map[int]int)
    v2 -= -9223372036854775808
    v4 := GenS1{F0: -1}
    if (0 != 1) {
        v1[2147483647] = "hello"
        v2 = v2
    }
    var v5 bool = (true && false)
    fmt.Println(v2, v5)
    var v6 float64 = (0.1 / (1.5 / (3.14159 + 0.0)))
    fmt.Println(v2)
    fmt.Println(v2)
    v5 = (!v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
