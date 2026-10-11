package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0(p0 uint32, p1 int8) bool {
    p1 += p1
    v6 := []int{1, 236, 729}
    _ = p0
    _ = p1
    _ = v6
    return true
}

func main() {
    var v1 int16 = 0
    v2 := []int{}
    v3 := make(map[string]string)
    v4 := make(map[uint32]int)
    var v5 uint = 18446744073709551615
    fmt.Println(v5, v5)
    v1 = v1
    fmt.Println(v5, v5)
    var v7 uint64 = 1
    v8 := GenS1{F0: ("" + "\n\t")}
    v9 := GenS1{F0: ("a" + "")}
    fmt.Println(v5, v7, v7)
    if true {
        fmt.Println(v5, v5, v1)
    } else {
        var v10 int8 = -1
        v11 := GenS1{F0: "hello"}
        var v12 uint64 = v7
        _ = v10
        _ = v11
        _ = v12
    }
    v7 = (v7 | 42)
    for fi13 := 0; fi13 < 17; fi13++ {
        fmt.Println(v1, fi13, v1)
        _ = fi13
    }
    v1 -= v1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
