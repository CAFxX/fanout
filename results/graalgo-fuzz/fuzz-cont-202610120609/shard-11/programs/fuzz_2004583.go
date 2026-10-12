package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }

func main() {
    var v1 bool = (-2147483648 >= 2147483647)
    var v2 int = 9223372036854775807
    var v3 bool = (-1 != -128)
    var v4 int8 = 53
    v5 := make(map[string]string)
    v6 := []string{}
    for fi7 := 0; fi7 < 12; fi7++ {
        var v8 uint = 42
        _ = len(v5)
        _ = fi7
        _ = v8
    }
    fmt.Println(v3, v1, v3)
    v3 = (1.5 <= 0.1)
    if (!v3) {
        var v9 float32 = 0.0
        _ = v9
    }
    v3 = ((float32(1.5) + -2.25) != 0.1)
    fmt.Println(v4, v3)
    fmt.Println(v1)
    var v10 float64 = 0.1
    v11 := v5[""]
    v12 := GenS1{F0: 165, F1: 1}
    if v3 {
        var v13 bool = (2 > 32767)
        var v14 int8 = 0
        _ = v13
        _ = v14
    } else {
        v3 = v1
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
