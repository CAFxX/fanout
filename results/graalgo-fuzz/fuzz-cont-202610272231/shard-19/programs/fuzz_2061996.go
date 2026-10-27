package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 uint16 }
func genFn0(p0 uint8, p1 uint, p2 string) int {
    if false {
        fmt.Println(p0, p1, p2)
    }
    v7 := GenS1{F0: -128, F1: 65535, F2: 1}
    p2 = p2
    for fi8 := 0; fi8 < 3; fi8++ {
        p2 = (p2 + (p2 + p2))
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 1
}

func main() {
    v1 := []string{"foo bar", "a"}
    var v2 int32 = 477
    var v3 float32 = float32((1e10 + 1e10))
    var v4 uint16 = 42
    v5 := make(map[int]uint32)
    v6 := make(map[string]string)
    if ((v4 >= v4) && true) {
        v2 -= 659
    }
    _ = len(v6)
    v4 = v4
    _ = len(v1)
    fmt.Println(v2)
    for fi9 := 0; fi9 < 3; fi9++ {
        fmt.Println(fi9, v3)
        _ = fi9
    }
    var v10 int = 32768
    var v11 uint8 = 2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
