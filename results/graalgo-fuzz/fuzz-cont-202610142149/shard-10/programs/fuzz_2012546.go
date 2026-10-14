package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0() string {
    for fi6 := 0; fi6 < 10; fi6++ {
        fmt.Println(fi6)
        fmt.Println(fi6)
        var v7 float64 = 0.1
        _ = fi6
        _ = v7
    }
    var v8 string = (("\n\t" + "foo bar") + "")
    v9 := GenS1{F0: 51}
    var v10 string = v8
    _ = v10
    _ = v8
    _ = v9
    return v8
}

func main() {
    var v1 uint64 = 1
    var v2 uint = 611
    v3 := make(map[string]bool)
    v4 := make(map[int]int)
    v5 := make(map[int]bool)
    var v11 uint8 = 255
    var v12 int16 = -32768
    v13 := GenS1{F0: (v11 | (v11 * v11))}
    var v14 int8 = 1
    v15 := []int{453, 2147483647, 0, 42}
    v1 += 0
    v2 -= uint(v1)
    fmt.Println(v11, v1, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
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
