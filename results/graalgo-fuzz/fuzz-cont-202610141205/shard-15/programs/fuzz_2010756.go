package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 bool }
type GenS2 struct { F0 uint16; F1 int }
func genFn0(p0 int32, p1 uint64) int8 {
    for fi5 := 0; fi5 < 17; fi5++ {
        var v6 string = "foo bar"
        v7 := GenS1{F0: 0, F1: fi5, F2: (-1 > int8(fi5))}
        _ = fi5
        _ = v6
        _ = v7
    }
    fmt.Println(p0)
    if true {
        v8 := []int{127, 244, 4294967295, -1}
        _ = v8
    }
    _ = p0
    _ = p1
    return 42
}
func genFn1(p0 int, p1 bool, p2 int64) int16 {
    switch "foo bar" {
    case "hello":
        p2 = -19
    case "a":
        switch ("" + ("\n\t" + ("\n\t" + "a"))) {
        case "\n\t":
            fmt.Println(p2, p0)
        default:
            fmt.Println(p1)
        }
    case "":
        var v9 uint64 = 551
        _ = v9
    default:
        fmt.Println(p0, p1)
    }
    p1 = false
    _ = p0
    _ = p1
    _ = p2
    return 1
}

func main() {
    var v1 int16 = -434
    var v2 int8 = -123
    v3 := make(map[int]string)
    var v4 int32 = -2147483648
    var v10 int64 = 42
    v10 = v10
    v4 = (v4 * v4)
    v10 -= v10
    var v11 int32 = 0
    var v12 float64 = 0.0
    var v13 int32 = v4
    v14 := []int{127}
    v15 := GenS2{F0: 262, F1: 1}
    _ = len(v3)
    _ = len(v14)
    var v16 float64 = (v12 + v12)
    v12 -= (3.14159 / v16)
    v4 -= v11
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
