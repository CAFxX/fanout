package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 uint64 }
type GenS2 struct { F0 int8; F1 string; F2 uint16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 bool }
func genFn0(p0 int32) uint {
    var v7 bool = false
    var v8 uint64 = 2
    _ = p0
    _ = v7
    _ = v8
    return 18446744073709551615
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[uint32]int)
    var v3 uint32 = 462
    v4 := make(map[uint32]uint32)
    var v5 string = ((("a" + ("a" + "a")) + "a") + "hello")
    v6 := []string{"a", v5}
    v9 := GenS1{F0: 4294967295, F1: 0, F2: 1}
    var v10 int = -972
    v11 := GenS2{F0: 100, F1: "hello", F2: 1}
    var v12 int8 = -128
    switch (v10 % ((v10) | 1)) {
    case -714:
        var v13 int16 = 42
        _ = v13
    default:
        switch v5 {
        case "\n\t":
            v6 = append(v6, "hello")
        }
    }
    if false {
        v14 := GenS3{F0: v3}
        v15 := GenS4{F0: true}
        var v16 uint64 = 909
        _ = v14
        _ = v15
        _ = v16
    }
    var v17 int32 = -1
    var v18 uint = 42
    v19 := GenS2{F0: v12, F1: (v5 + ("a" + "\n\t")), F2: 32767}
    var v20 uint = (v18 + v18)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
