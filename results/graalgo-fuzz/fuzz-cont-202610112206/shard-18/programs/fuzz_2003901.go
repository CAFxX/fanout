package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8; F1 int32 }

func main() {
    var v1 float64 = float64(1e10)
    var v2 uint8 = 255
    v3 := []int{30}
    var v4 string = "a"
    var v5 float64 = v1
    v6 := make(map[string]bool)
    v7 := GenS1{F0: 1}
    fmt.Println(v4)
    switch 42 {
    case 2147483647:
        var v8 uint = 255
        _ = v8
    case -347:
        v9 := GenS2{F0: 100, F1: 543}
        _ = v9
    case 32767:
        var v10 bool = (1e10 != float32((float64(3.14159) / v5)))
        _ = v10
    }
    fmt.Println(v2, v4)
    var v11 int64 = 9223372036854775807
    var v12 string = "hello"
    var v13 uint16 = 2
    if (42 != uint64(v2)) {
        v12 = ""
    }
    v14 := v6[(v12 + "\n\t")]
    v4 = ""
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
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
}
