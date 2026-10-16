package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    var v1 bool = true
    v2 := []int{-1, 128, 127}
    v3 := make(map[uint32]int)
    v4 := make(map[string]bool)
    var v5 string = (("foo bar" + "foo bar") + ("" + ""))
    var v6 int32 = 1
    fmt.Println(v1)
    v7 := GenS1{F0: -32768}
    v2 = append(v2, 220)
    var v8 int32 = (^(v6))
    fmt.Println(v5)
    switch v5 {
    case "a":
        v5 = (v5 + "a")
    default:
        v5 = "\n\t"
    }
    fmt.Println(v8, v5, v6)
    if ((v6 > v6) || (-32768 > 32767)) {
        v9 := GenS1{F0: 1}
        _ = v9
    } else {
        v5 = v5
        if len(v2) > 0 {
            v2[0] = 2147483647
        }
        v1 = true
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
