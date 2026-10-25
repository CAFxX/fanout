package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]bool)
    v3 := make(map[string]string)
    var v4 string = ""
    v5 := make(map[string]int)
    var v6 int16 = 1000
    switch -1 {
    case -9223372036854775808:
        var v7 uint = 42
        _ = v7
    case -880:
        fmt.Println(v4, v6)
    default:
        v8 := GenS1{F0: true}
        _ = v8
    }
    v9 := []int{}
    fmt.Println(v4, v6)
    if len(v9) >= 2 {
        v9 = v9[1:len(v9)]
    }
    switch v4 {
    case "hello":
        v10 := GenS1{F0: false}
        _ = v10
    case "\n\t":
        var v11 int32 = -2147483648
        _ = v11
    default:
        var v12 bool = true
        _ = v12
    }
    var v13 int = -2147483648
    fmt.Println(v4)
    if len(v9) >= 2 {
        v9 = v9[1:len(v9)]
    }
    v4 = v4
    v14 := GenS1{F0: true}
    var v15 float64 = float64(-2.25)
    _ = len(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
