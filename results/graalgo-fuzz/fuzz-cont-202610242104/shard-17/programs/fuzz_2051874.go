package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32 }
func genFn0(p0 uint64) int {
    p0 = p0
    v4 := GenS1{F0: (("foo bar" + "a") + "foo bar")}
    p0 = p0
    p0 = (p0 / ((p0) | 1))
    v5 := GenS1{F0: "\n\t"}
    _ = p0
    _ = v4
    _ = v5
    return 127
}

func main() {
    v1 := make(map[string]int)
    var v2 bool = (((true || false) || (3.14159 >= 1e10)) || false)
    v3 := make(map[uint32]bool)
    var v6 bool = (1 >= -101)
    var v7 bool = false
    var v8 int32 = -1
    v9 := []int{2147483647, genFn0(2), 42, genFn0(845)}
    for fi10 := 0; fi10 < 6; fi10++ {
        v7 = v7
        _ = fi10
    }
    var v11 int64 = 0
    v12 := GenS2{F0: v8}
    v8 = (^(v8))
    if len(v9) > 0 {
        v9[0] = 9223372036854775807
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
