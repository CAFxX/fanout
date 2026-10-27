package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0(p0 int64, p1 bool) uint16 {
    v6 := []int{}
    var v7 uint = 65535
    if len(v6) > 0 {
        v6[0] = 32767
    }
    p0 += p0
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 808
}

func main() {
    var v1 int8 = -128
    v2 := make(map[string]uint32)
    var v3 string = (("hello" + "\n\t") + "\n\t")
    v4 := make(map[uint32]bool)
    var v5 uint = 42
    fmt.Println(v1)
    fmt.Println(v1, v3)
    var v8 int8 = v1
    fmt.Println(v5, v8)
    _ = len(v2)
    var v9 uint64 = 4294967295
    v10 := GenS1{F0: v1}
    v5 -= v5
    v11 := []int{308}
    var v12 uint32 = 0
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
