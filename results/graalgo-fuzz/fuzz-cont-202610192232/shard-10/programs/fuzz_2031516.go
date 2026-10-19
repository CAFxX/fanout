package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 uint16 }

func main() {
    var v1 uint = 42
    v2 := make(map[uint32]int)
    var v3 float64 = 1e10
    v4 := []uint32{42, 1, 42, 288}
    v5 := []string{("hello" + "")}
    var v6 float32 = 1.5
    var v7 uint64 = 18446744073709551615
    if (9223372036854775807 == 414) {
        v1 = v1
        fmt.Println(v6)
        if false {
            fmt.Println(v1, v1, v6)
        } else {
            _ = len(v2)
            var v8 float64 = 1e10
            _ = v8
        }
    }
    _ = len(v2)
    if len(v5) > 0 {
        v5[0] = ("foo bar" + "\n\t")
    }
    v9 := GenS1{F0: (v7 ^ 4294967295), F1: 223, F2: 1}
    for fi10 := 0; fi10 < 8; fi10++ {
        v11 := GenS1{F0: (^(v7)), F1: (fi10 * int(v7)), F2: 42}
        var v12 uint64 = 0
        _ = fi10
        _ = v11
        _ = v12
    }
    v3 = 1e10
    for fi13 := 0; fi13 < 19; fi13++ {
        v7 = 42
        _ = fi13
    }
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
