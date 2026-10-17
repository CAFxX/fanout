package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool }

func main() {
    var v1 uint64 = 42
    var v2 int8 = 48
    var v3 float32 = 0.1
    v4 := make(map[string]int)
    var v5 int64 = 42
    var v6 float32 = v3
    var v7 float64 = ((0.1 / float64(v3)) + 1e10)
    switch ("foo bar" + "\n\t") {
    case "":
        v8 := []int{32767, 42}
        _ = v8
    case "\n\t":
        fmt.Println(v1)
    case "a":
        var v9 uint64 = v1
        _ = v9
    default:
        v6 -= (1e10 + 0.1)
    }
    fmt.Println(v3)
    v10 := GenS1{F0: 1}
    v5 = v5
    v11 := []int{}
    v2 += v2
    v7 = ((-2.25 - (0.1 / v7)) + float64(((1e10 / (v6 / (v6 - float32(v7)))) - v3)))
    if len(v11) > 0 {
        v11[0] = 0
    }
    v12 := GenS2{F0: false}
    var v13 int16 = -32768
    v14 := GenS2{F0: (((42 < v2) && false) && (!false))}
    var v15 string = (("hello" + "") + ("a" + "hello"))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
