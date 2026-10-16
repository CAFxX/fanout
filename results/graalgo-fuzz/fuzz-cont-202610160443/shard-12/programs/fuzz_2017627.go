package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    var v1 uint64 = 2
    var v2 int32 = 347
    v3 := []int{9223372036854775807, -305, 4294967295}
    var v4 float32 = 3.14159
    v5 := make(map[string]uint32)
    v6 := GenS1{F0: 515}
    _ = len(v5)
    var v7 int16 = 976
    v2 = (v2 & 42)
    var v8 int32 = v2
    var v9 uint16 = 890
    v10 := GenS1{F0: 0}
    v11 := v5[(("hello" + "a") + ("hello" + "foo bar"))]
    var v12 uint16 = (^(v9))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
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
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
