package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0(p0 uint8, p1 bool) uint8 {
    p0 = p0
    p1 = true
    _ = p0
    _ = p1
    return (^(p0))
}
func genFn1(p0 int16) string {
    var v6 string = "a"
    var v7 uint = 18446744073709551615
    v8 := GenS1{F0: -47}
    v9 := []int{9223372036854775807}
    var v10 bool = true
    _ = p0
    _ = v10
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return v6
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]bool)
    var v3 uint16 = 32767
    v4 := []int{127, -1, 794}
    var v5 int64 = -1
    var v11 float64 = 1e10
    fmt.Println(v11, v5)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v12 uint8 = 1
    v13 := GenS1{F0: 42}
    v14 := GenS1{F0: 100}
    v15 := GenS1{F0: 75}
    _ = len(v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
