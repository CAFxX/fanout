package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
type GenS2 struct { F0 uint32; F1 uint32; F2 uint32 }
func genFn0(p0 uint, p1 uint16) int32 {
    p1 = p1
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    return 2147483647
}

func main() {
    v1 := make(map[string]bool)
    var v2 bool = true
    var v3 uint64 = 1
    v4 := []string{}
    var v5 int = 128
    v6 := GenS1{F0: 4294967295, F1: 1}
    fmt.Println(v5, v3, v2)
    v3 += v3
    v3 -= (^(v3))
    v4 = append(v4, "\n\t")
    v7 := v1["\n\t"]
    for fi8 := 0; fi8 < 20; fi8++ {
        delete(v1, "hello")
        v9 := GenS1{F0: 4294967295, F1: 127}
        _ = fi8
        _ = v9
    }
    if ((-(v5)) < v5) {
        v10 := GenS2{F0: 396, F1: 106, F2: 2147483647}
        _ = v10
    }
    fmt.Println(v3, v2)
    var v11 int8 = int8(v5)
    var v12 float32 = 1.5
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
