package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16 }
func genFn0(p0 int16) int16 {
    p0 += (p0 * p0)
    v5 := GenS1{F0: (-39 == 127)}
    v6 := GenS2{F0: p0}
    p0 -= p0
    _ = p0
    _ = v5
    _ = v6
    return p0
}

func main() {
    var v1 float32 = 1.5
    var v2 float64 = (0.1 - 1.5)
    var v3 int64 = -9223372036854775808
    v4 := []int{-2147483648, 2147483647, -44, -261}
    var v7 string = "\n\t"
    _ = len(v4)
    var v8 uint8 = 226
    v9 := GenS2{F0: -32768}
    var v10 int = 717
    if true {
        for fi11 := 0; fi11 < 9; fi11++ {
            var v12 uint32 = 354
            _ = len(v4)
            _ = fi11
            _ = v12
        }
        var v13 int16 = genFn0(1)
        _ = v13
    }
    v2 -= v2
    v10 = v10
    v14 := GenS1{F0: true}
    switch -895 {
    case -290:
        var v15 float32 = 1.5
        _ = v15
    }
    fmt.Println(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
