package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 bool; F1 string; F2 int16 }

func main() {
    v1 := []string{"", "hello"}
    v2 := make(map[string]string)
    v3 := make(map[string]bool)
    v4 := make(map[string]int)
    var v5 uint64 = 18446744073709551615
    v6 := make(map[uint32]string)
    v5 -= v5
    v7 := GenS1{F0: 42}
    if false {
        v8 := GenS2{F0: false, F1: "", F2: -913}
        var v9 int32 = 1
        v10 := GenS2{F0: (-594 > 0), F1: ("\n\t" + "\n\t"), F2: 1000}
        _ = v10
        _ = v8
        _ = v9
    } else {
        fmt.Println(v5)
        var v11 float32 = 0.1
        switch 42 {
        case -307:
            var v12 uint = 65535
            _ = v12
        default:
            v13 := GenS2{F0: (!false), F1: "hello", F2: 32767}
            _ = v13
        }
        _ = v11
    }
    fmt.Println(v5)
    var v14 int16 = -576
    fmt.Println(v5)
    fmt.Println(v14, v14)
    v15 := GenS2{F0: (true && false), F1: "a", F2: (v14 - v14)}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
