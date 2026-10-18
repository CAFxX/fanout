package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
func genFn0() string {
    var v4 bool = (!(-889 != 0))
    v5 := []int{-9223372036854775808}
    var v6 uint8 = 2
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    _ = v4
    _ = v5
    _ = v6
    return "foo bar"
}

func main() {
    v1 := make(map[int]bool)
    var v2 int32 = -1
    var v3 int16 = -32768
    var v7 uint = 2
    v2 = 892
    v8 := GenS1{F0: 0, F1: 32768}
    var v9 int = 42
    var v10 uint64 = 18446744073709551615
    var v11 float32 = 0.0
    v9 += int(v10)
    v12 := GenS1{F0: v7, F1: int(v3)}
    var v13 bool = (-1 >= 127)
    fmt.Println(v3)
    var v14 uint8 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
