package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 int8 }

func main() {
    var v1 int16 = -204
    var v2 uint32 = 2147483648
    v3 := []string{"a", ("\n\t" + ("a" + "\n\t")), "foo bar"}
    var v4 uint64 = 0
    v5 := make(map[int]bool)
    var v6 bool = (1 != 0)
    v7 := v5[1]
    v8 := GenS1{F0: int32(v4), F1: 139}
    v9 := GenS1{F0: -1, F1: 32767}
    v6 = false
    v10 := GenS2{F0: -1}
    var v11 int32 = 2147483647
    delete(v5, int(v1))
    var v12 bool = (!((v6 || false) || (-2 == 42)))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
