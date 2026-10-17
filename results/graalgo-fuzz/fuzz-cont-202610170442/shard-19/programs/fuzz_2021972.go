package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    v1 := []string{"\n\t", "foo bar"}
    var v2 string = "hello"
    v3 := []uint32{2147483647}
    var v4 uint32 = 65535
    v5 := make(map[uint32]bool)
    v6 := GenS1{F0: 324}
    v4 -= 696
    if (1 <= uint(v4)) {
        var v7 int = -870
        _ = v7
    }
    var v8 int = 127
    var v9 uint32 = 0
    v8 = (v8 << ((65535) & 63))
    var v10 uint64 = 18446744073709551615
    if false {
        v9 = (v4 >> ((100) & 63))
    } else {
        if len(v3) > 0 {
            v3[0] = v4
        }
    }
    v11 := GenS1{F0: -176}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
