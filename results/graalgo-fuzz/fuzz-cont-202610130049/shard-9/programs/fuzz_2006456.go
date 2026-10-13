package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 uint }
type GenS2 struct { F0 int; F1 int32 }

func main() {
    v1 := make(map[int]int)
    var v2 uint = 4294967295
    v3 := make(map[string]int)
    var v4 string = ""
    v5 := GenS1{F0: 2147483648, F1: 548, F2: 255}
    fmt.Println(v2)
    var v6 float64 = 1.5
    fmt.Println(v4, v2, v6)
    switch 42 {
    case 42:
        v7 := []int{151, 32767, 0, -272}
        _ = v7
    default:
        var v8 int32 = 2147483647
        _ = v8
    }
    fmt.Println(v2)
    var v9 bool = (!(true && false))
    v10 := []int{}
    var v11 bool = (0.0 >= 0.0)
    var v12 int8 = -60
    switch v4 {
    case "\n\t":
        v13 := GenS2{F0: -1, F1: -901}
        _ = v13
    case "a":
        var v14 uint32 = 0
        _ = v14
    case "":
        fmt.Println(v12, v4)
    default:
        v15 := GenS2{F0: 127, F1: -237}
        _ = v15
    }
    if v9 {
        fmt.Println(v4, v11, v6)
    } else {
        if true {
            v12 = v12
            v2 -= 2
        }
        v12 = 100
        for fi16 := 0; fi16 < 8; fi16++ {
            delete(v3, (v4 + "foo bar"))
            _ = fi16
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
