package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int64 }
func genFn0() int64 {
    for fi7 := 0; fi7 < 1; fi7++ {
        var v8 float32 = float32((-2.25 - ((3.14159 / (0.0 - (1e10 / -2.25))) + 1.5)))
        fmt.Println(fi7)
        if (false && true) {
            fmt.Println(v8, fi7)
        }
        _ = fi7
        _ = v8
    }
    v9 := []int{1, 42, -385, 127}
    var v10 int8 = 1
    v10 = 21
    _ = v10
    _ = v9
    return -9223372036854775808
}
func genFn1(p0 uint16, p1 uint16, p2 bool) int64 {
    p2 = (6 < genFn0())
    var v11 int8 = 100
    switch 806 {
    case 32768:
        p0 = 204
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return 0
}

func main() {
    v1 := make(map[int]string)
    var v2 uint16 = 2
    var v3 uint8 = 173
    var v4 int16 = 42
    v5 := []string{}
    v6 := make(map[uint32]string)
    if true {
        v12 := GenS1{F0: -95}
        v13 := GenS2{F0: 971}
        _ = v12
        _ = v13
    }
    var v14 int = -2147483648
    var v15 uint8 = 255
    v16 := v6[4294967295]
    v17 := GenS1{F0: -128}
    v15 -= v15
    fmt.Println(v2)
    var v18 int8 = 127
    fmt.Println(v18)
    fmt.Println(v4, v2, v14)
    delete(v6, 2147483647)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
