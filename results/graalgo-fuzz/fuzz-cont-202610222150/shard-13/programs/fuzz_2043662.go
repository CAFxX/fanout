package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16 }
func genFn0(p0 int64) string {
    var v6 int = 128
    v7 := []int{}
    var v8 int64 = -9223372036854775808
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return "hello"
}
func genFn1(p0 int32, p1 int32) uint32 {
    p1 = p0
    p1 = p1
    p0 = (p1 ^ 42)
    _ = p0
    _ = p1
    return 0
}

func main() {
    v1 := make(map[string]bool)
    var v2 float32 = (((float32((3.14159 * (3.14159 + float64(0.1)))) / 1.0) / float32((1e10 * ((1e10 / 1.5) / (0.1 * 1.5))))) - 0.0)
    v3 := []uint32{1, 476, 4294967295, 919}
    v4 := []int{4294967295, -1, 32767}
    v5 := make(map[string]int)
    var v9 bool = (false && (true || false))
    v2 -= float32(-2.25)
    v9 = true
    var v10 int64 = -1
    v11 := GenS1{F0: 178}
    switch 4294967295 {
    case -217:
        var v12 int8 = -128
        _ = v12
    }
    v2 += (float32((3.14159 - -2.25)) * v2)
    var v13 uint = 18446744073709551615
    switch "\n\t" {
    case "a":
        fmt.Println(v2, v9)
    }
    v14 := GenS2{F0: 2}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
