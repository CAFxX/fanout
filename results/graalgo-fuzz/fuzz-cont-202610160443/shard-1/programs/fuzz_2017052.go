package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }

func main() {
    v1 := make(map[uint32]string)
    var v2 uint = 1
    v3 := make(map[int]uint32)
    v4 := []string{(("" + ("" + "\n\t")) + "a"), "foo bar"}
    var v5 uint = (^(v2))
    var v6 uint = 0
    v2 += (v6 | (v6 & v6))
    for fi7 := 0; fi7 < 19; fi7++ {
        v8 := GenS1{F0: (4294967295 != 0), F1: 127}
        var v9 string = "hello"
        _ = fi7
        _ = v8
        _ = v9
    }
    v5 = v2
    v4 = append(v4, "hello")
    var v10 uint64 = 4294967295
    var v11 uint16 = 42
    for fi12 := 0; fi12 < 8; fi12++ {
        var v13 uint64 = 0
        var v14 int32 = -560
        if (false && (true && true)) {
            fmt.Println(v10, v14, v5)
            var v15 int = (^(fi12))
            _ = v15
        }
        _ = fi12
        _ = v13
        _ = v14
    }
    if false {
        var v16 uint64 = (^(v10))
        _ = v16
    }
    v10 -= uint64(v11)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
