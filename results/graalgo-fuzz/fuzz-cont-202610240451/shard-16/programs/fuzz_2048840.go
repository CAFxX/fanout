package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 int64 }

func main() {
    v1 := []uint32{65535, 0}
    v2 := []string{(("a" + ((("\n\t" + ("a" + "\n\t")) + "\n\t") + "\n\t")) + "a"), ("\n\t" + "\n\t")}
    var v3 float32 = 0.0
    var v4 uint16 = 32767
    v5 := make(map[int]int)
    v6 := make(map[uint32]bool)
    var v7 int64 = 1
    fmt.Println(v3, v4, v7)
    fmt.Println(v4, v3, v4)
    var v8 uint64 = 0
    v8 += v8
    var v9 int16 = 42
    if false {
        v10 := GenS1{F0: "\n\t", F1: 0, F2: -1}
        if false {
            v11 := GenS1{F0: "a", F1: 691, F2: 1}
            fmt.Println(v7, v3)
            v4 -= 575
            _ = v11
        }
        _ = v10
    }
    fmt.Println(v4, v3, v4)
    var v12 bool = (v8 >= 0)
    switch 4294967295 {
    case 42:
        var v13 bool = false
        _ = v13
    }
    v12 = (v9 < v9)
    delete(v6, 885)
    v9 = v9
    _ = len(v6)
    v14 := v5[9223372036854775807]
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
