package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0() int8 {
    v6 := []int{727, 127, -2147483648, 128}
    v7 := GenS1{F0: 1}
    if false {
        v6 = append(v6, 32768)
        var v8 float64 = ((0.0 * 1e10) - 0.0)
        _ = v8
    } else {
        var v9 int64 = 664
        if len(v6) > 0 {
            v6[0] = 42
        }
        _ = v9
    }
    _ = len(v6)
    _ = v6
    _ = v7
    return -62
}

func main() {
    var v1 uint8 = 64
    v2 := []uint32{}
    var v3 float64 = 1e10
    v4 := make(map[int]bool)
    v5 := []uint32{409}
    var v10 bool = false
    v10 = ((v10 && v10) || false)
    v10 = (42 < 58)
    var v11 string = "\n\t"
    var v12 string = ((v11 + v11) + v11)
    var v13 string = (v11 + (v11 + v12))
    v2 = append(v2, 65535)
    if (v10 && v10) {
        if false {
            fmt.Println(v10)
        }
    }
    var v14 uint = 780
    v4[9223372036854775807] = (-1 < -294)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
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
}
