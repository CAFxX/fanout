package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0() string {
    v4 := []int{-758, -2147483648, 42, 128}
    var v5 int64 = 42
    for fi6 := 0; fi6 < 12; fi6++ {
        switch "foo bar" {
        case "":
            v5 += v5
        default:
            var v7 int16 = 42
            _ = v7
        }
        var v8 int64 = v5
        v8 -= v8
        _ = fi6
        _ = v8
    }
    _ = v4
    _ = v5
    return ("foo bar" + "hello")
}

func main() {
    var v1 int64 = -497
    var v2 bool = true
    v3 := make(map[int]int)
    var v9 float64 = 1e10
    v10 := GenS1{F0: v1}
    delete(v3, -9223372036854775808)
    v9 = (-2.25 * v9)
    v1 += 1
    var v11 int = -1
    v9 -= 0.1
    v3[(v11 ^ v11)] = v11
    v11 += int(v1)
    v1 = (v1 ^ v1)
    for fi12 := 0; fi12 < 1; fi12++ {
        var v13 int64 = v1
        _ = fi12
        _ = v13
    }
    var v14 uint = 718
    v15 := []int{v11, v11, v11}
    var v16 float32 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v9
}
