package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }

func main() {
    var v1 float64 = ((0.0 * 0.0) - 3.14159)
    v2 := make(map[int]bool)
    var v3 uint8 = 0
    var v4 string = "a"
    v5 := make(map[int]uint32)
    v3 -= (^(v3))
    var v6 string = v4
    fmt.Println(v1)
    v1 = (float64(1.5) + v1)
    v1 -= float64(0.1)
    if true {
        _ = len(v2)
        v3 = v3
    }
    var v7 int16 = 1
    if ((v3 % ((v3) | 1)) < v3) {
        fmt.Println(v3, v1)
        var v8 int32 = 2147483647
        switch "\n\t" {
        case "":
            v1 += v1
        case "a":
            v9 := []int{}
            _ = v9
        default:
            var v10 float32 = 3.14159
            _ = v10
        }
        _ = v8
    } else {
        for fi11 := 0; fi11 < 9; fi11++ {
            v3 = (^(v3))
            _ = fi11
        }
        fmt.Println(v4)
    }
    if (0 != 42) {
        v6 = ("a" + (v4 + ""))
        v4 = ("" + "a")
        v12 := []int{1}
        _ = v12
    }
    v13 := GenS1{F0: 9223372036854775807, F1: 191}
    var v14 string = (v6 + v6)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
