package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 string }

func main() {
    v1 := make(map[string]bool)
    var v2 string = "\n\t"
    v3 := make(map[string]uint32)
    var v4 float32 = 1.5
    v5 := make(map[string]int)
    var v6 uint = 61
    v6 -= (v6 | 0)
    switch 128 {
    case -925:
        fmt.Println(v4)
    }
    for fi7 := 0; fi7 < 17; fi7++ {
        var v8 uint32 = 4294967295
        _ = fi7
        _ = v8
    }
    v4 = float32(1e10)
    switch -2147483648 {
    case 590:
        fmt.Println(v4, v4)
    case 32767:
        v6 -= v6
    default:
        _ = len(v1)
    }
    switch "foo bar" {
    case "hello":
        switch 32767 {
        case 4294967295:
            v6 -= (v6 | (v6 % ((42) | 1)))
        case -1:
            fmt.Println(v4, v2)
        default:
            v9 := []int{2147483647, 2147483647}
            _ = v9
        }
    case "a":
        if false {
            var v10 int8 = -1
            _ = len(v3)
            fmt.Println(v4, v4, v6)
            _ = v10
        }
    }
    var v11 float64 = -2.25
    fmt.Println(v4, v4, v11)
    fmt.Println(v11, v6)
    v1["hello"] = (2 < 0)
    fmt.Println(v4, v11)
    var v12 string = v2
    for fi13 := 0; fi13 < 3; fi13++ {
        v14 := GenS1{F0: -387, F1: "foo bar", F2: v12}
        fmt.Println(fi13, fi13, v6)
        var v15 int8 = 1
        _ = fi13
        _ = v14
        _ = v15
    }
    v6 += (^(v6))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
