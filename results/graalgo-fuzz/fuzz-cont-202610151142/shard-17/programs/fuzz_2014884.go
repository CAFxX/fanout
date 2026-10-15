package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    var v1 string = ("hello" + ("a" + ""))
    v2 := make(map[string]uint32)
    var v3 uint32 = 65535
    v4 := []int{}
    v5 := []string{v1, v1, v1}
    var v6 uint = 4294967295
    var v7 uint8 = 1
    var v8 uint32 = 1
    v9 := GenS1{F0: 291}
    v8 = 2
    var v10 uint64 = 4294967295
    fmt.Println(v6, v8, v6)
    var v11 uint16 = 494
    var v12 uint64 = (v10 / ((v10) | 1))
    var v13 bool = (!(65535 != uint32(v6)))
    v12 = v12
    var v14 uint = v6
    v7 -= v7
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
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
