package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
func genFn0(p0 bool, p1 uint8, p2 int64) int8 {
    v7 := GenS1{F0: -483, F1: -906}
    v8 := GenS1{F0: 1, F1: 197}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 1
}

func main() {
    var v1 int16 = -1
    var v2 int = 4294967295
    var v3 uint64 = 18446744073709551615
    v4 := []string{}
    v5 := make(map[int]uint32)
    var v6 int64 = -741
    fmt.Println(v6)
    var v9 uint32 = 65535
    _ = len(v5)
    v6 += (v6 | v6)
    v6 += v6
    v4 = append(v4, ("a" + ("\n\t" + ("" + (("a" + ("hello" + "hello")) + "foo bar")))))
    v10 := v5[-85]
    v11 := GenS1{F0: 2147483647, F1: v6}
    v12 := GenS1{F0: -1, F1: (-(v6))}
    _ = len(v5)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v10 = (^(v10))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
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
    _ = v9
}
