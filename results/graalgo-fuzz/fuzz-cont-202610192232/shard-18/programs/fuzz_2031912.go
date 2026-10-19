package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0(p0 int64, p1 uint8, p2 string) int32 {
    v4 := []int{-1}
    p0 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return -2147483648
}

func main() {
    var v1 string = "hello"
    v2 := make(map[uint32]int)
    v3 := make(map[uint32]bool)
    v5 := GenS1{F0: 32767}
    var v6 int64 = -9223372036854775808
    var v7 float32 = float32(0.1)
    var v8 uint = 4294967295
    var v9 uint16 = 886
    v10 := []int{-666, 127}
    fmt.Println(v7, v9, v8)
    _ = len(v2)
    if (!(236 == 127)) {
        v8 = v8
        var v11 uint16 = v9
        _ = v11
    }
    for fi12 := 0; fi12 < 19; fi12++ {
        var v13 int = fi12
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 5; fi14++ {
        _ = len(v10)
        _ = fi14
    }
    fmt.Println(v6, v6)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
