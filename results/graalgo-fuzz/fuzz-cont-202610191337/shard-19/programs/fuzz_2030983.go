package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 int }

func main() {
    v1 := make(map[string]int)
    var v2 uint64 = 838
    v3 := make(map[string]string)
    v4 := make(map[string]bool)
    switch "hello" {
    case "a":
        v5 := []int{995, -1, 339, -618}
        _ = v5
    default:
        var v6 string = ("a" + ("a" + "\n\t"))
        _ = v6
    }
    var v7 uint32 = 2147483648
    var v8 int = -1
    v8 = v8
    for fi9 := 0; fi9 < 1; fi9++ {
        fmt.Println(v2, fi9)
        v8 = -9223372036854775808
        _ = fi9
    }
    var v10 int8 = -1
    v11 := GenS1{F0: "\n\t", F1: v10, F2: v8}
    v7 = v7
    var v12 int16 = -32768
    v8 = v8
    for fi13 := 0; fi13 < 15; fi13++ {
        switch 1 {
        case -73:
            var v14 int16 = v12
            _ = v14
        default:
            var v15 int32 = 0
            _ = v15
        }
        var v16 int8 = (v10 * -77)
        _ = fi13
        _ = v16
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
