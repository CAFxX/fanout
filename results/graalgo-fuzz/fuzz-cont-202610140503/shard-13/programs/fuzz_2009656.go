package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool; F1 int8; F2 int }
func genFn0() bool {
    switch "\n\t" {
    case "a":
        var v4 string = "hello"
        _ = v4
    case "":
        v5 := []int{-984, 127}
        _ = v5
    default:
        var v6 uint16 = 65535
        _ = v6
    }
    var v7 string = "hello"
    _ = v7
    return false
}

func main() {
    var v1 float32 = 1e10
    v2 := make(map[int]bool)
    v3 := make(map[int]string)
    fmt.Println(v1)
    fmt.Println(v1)
    v8 := GenS1{F0: ("a" + "hello")}
    var v9 bool = true
    v9 = v9
    v1 -= float32(1.5)
    v9 = v9
    switch "" {
    case "hello":
        delete(v3, -2147483648)
    case "a":
        v10 := GenS2{F0: v9, F1: 100, F2: -9223372036854775808}
        _ = v10
    case "foo bar":
        for fi11 := 0; fi11 < 13; fi11++ {
            v1 -= (v1 + ((float32((float64(-2.25) * 0.1)) / ((((1.5 / 0.1) + ((3.14159 * float32(1.5)) * v1)) - 1e10) * 1e10)) / (float32((1e10 + -2.25)) - v1)))
            _ = fi11
        }
    }
    v9 = true
    var v12 int64 = 9223372036854775807
    _ = len(v3)
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
