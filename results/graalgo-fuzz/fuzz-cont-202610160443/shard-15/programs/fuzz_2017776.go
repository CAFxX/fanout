package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 uint64 }
type GenS2 struct { F0 int16 }

func main() {
    var v1 uint32 = 584
    v2 := make(map[uint32]string)
    var v3 uint = 65535
    var v4 uint32 = v1
    v5 := []string{("hello" + "\n\t"), ("a" + "foo bar"), "a", "a"}
    switch 127 {
    case -9223372036854775808:
        v1 += 743
    default:
        var v6 int64 = 0
        _ = v6
    }
    for fi7 := 0; fi7 < 1; fi7++ {
        var v8 uint32 = 466
        _ = fi7
        _ = v8
    }
    if false {
        v9 := GenS1{F0: false, F1: 1, F2: 4294967295}
        var v10 uint = v3
        _ = v10
        _ = v9
    } else {
        fmt.Println(v4, v3, v3)
        var v11 float64 = (-2.25 * (0.1 + (0.0 / 0.1)))
        _ = v11
    }
    var v12 string = "a"
    var v13 string = "foo bar"
    v12 = "a"
    var v14 int8 = 0
    var v15 int8 = v14
    fmt.Println(v4, v3)
    v12 = "hello"
    v4 = (v4 - (v4 | v1))
    v16 := GenS2{F0: 0}
    v17 := v2[2]
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
