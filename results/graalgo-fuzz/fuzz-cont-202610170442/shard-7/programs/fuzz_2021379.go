package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int }

func main() {
    var v1 float64 = float64(-2.25)
    var v2 uint16 = 940
    var v3 float32 = 0.1
    var v4 uint = 65535
    v5 := make(map[string]string)
    v6 := make(map[int]string)
    var v7 string = "foo bar"
    var v8 bool = (-2.25 >= v1)
    var v9 uint16 = (^(v2))
    if true {
        var v10 bool = (v8 && (((v7 + v7) <= v7) || v8))
        fmt.Println(v2)
        _ = v10
    }
    v11 := GenS1{F0: 58}
    switch (v7 + v7) {
    case "hello":
        v12 := []int{}
        _ = v12
    case "a":
        var v13 float32 = (v3 + ((float32((0.0 - -2.25)) + v3) * (0.1 / ((float32(1.5) * v3) + (float32(0.0) - (v3 * 0.1))))))
        _ = v13
    default:
        v4 = (v4 % ((v4) | 1))
    }
    if true {
        v14 := GenS2{F0: 9223372036854775807}
        v15 := GenS2{F0: 2147483647}
        _ = v14
        _ = v15
    }
    v16 := GenS1{F0: 127}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
