package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint }
func genFn0(p0 uint32) int64 {
    v7 := GenS1{F0: -57}
    var v8 uint64 = 18446744073709551615
    _ = p0
    _ = v7
    _ = v8
    return 0
}

func main() {
    v1 := make(map[uint32]int)
    var v2 bool = true
    v3 := make(map[uint32]string)
    var v4 int32 = -135
    v5 := make(map[uint32]int)
    var v6 float64 = 1e10
    v4 -= v4
    v2 = (v2 && v2)
    if true {
        switch -2147483648 {
        case 42:
            var v9 float64 = (float64(-2.25) * v6)
            _ = v9
        }
    }
    v10 := []int{-9223372036854775808, -49, -428}
    v11 := GenS1{F0: -384}
    var v12 uint16 = 0
    var v13 int16 = 1
    v13 = v13
    v14 := GenS2{F0: 0}
    v6 = 0.1
    fmt.Println(v13, v12)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
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
}
