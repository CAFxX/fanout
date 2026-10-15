package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 uint64 }
type GenS2 struct { F0 int64; F1 uint }
func genFn0(p0 int8, p1 uint8) int8 {
    fmt.Println(p1, p0)
    var v7 int16 = 0
    v8 := []int{573, 2147483647, 9223372036854775807, 639}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return p0
}
func genFn1() int64 {
    var v9 int64 = 1
    fmt.Println(v9)
    for fi10 := 0; fi10 < 12; fi10++ {
        var v11 bool = false
        _ = fi10
        _ = v11
    }
    _ = v9
    return (v9 | v9)
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint16 = 2
    var v3 string = "foo bar"
    v4 := make(map[string]string)
    var v5 float64 = ((float64(1e10) / 1.0) + 0.0)
    v6 := make(map[uint32]bool)
    v12 := v4[v3]
    v5 += v5
    fmt.Println(v3, v12, v2)
    var v13 bool = false
    v14 := v6[2147483648]
    v15 := GenS1{F0: 2, F1: -891, F2: 0}
    if (127 == -1) {
        fmt.Println(v2, v3)
        var v16 float64 = v5
        _ = v16
    }
    var v17 uint8 = 127
    var v18 bool = v14
    fmt.Println(v13)
    switch 127 {
    case 792:
        v19 := []int{574, 660}
        _ = v19
    case 2147483647:
        for fi20 := 0; fi20 < 4; fi20++ {
            fmt.Println(v2, v13, v3)
            _ = fi20
        }
    case 214:
        v13 = v13
    }
    v4["a"] = "a"
    v17 = v17
    v12 = "\n\t"
    v21 := GenS2{F0: 1, F1: 4294967295}
    v12 = v3
    v22 := GenS2{F0: -361, F1: 1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
