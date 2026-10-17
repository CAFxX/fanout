package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    v1 := make(map[string]string)
    var v2 uint32 = 2147483648
    var v3 int32 = 0
    v4 := GenS1{F0: (!(1e10 >= 0.0))}
    for fi5 := 0; fi5 < 5; fi5++ {
        var v6 string = "a"
        v7 := []int{(fi5 ^ fi5), fi5, fi5, fi5}
        v3 = v3
        _ = fi5
        _ = v6
        _ = v7
    }
    var v8 float64 = -2.25
    var v9 string = "a"
    _ = len(v1)
    for fi10 := 0; fi10 < 15; fi10++ {
        v3 -= -1
        var v11 int8 = int8(v2)
        _ = fi10
        _ = v11
    }
    var v12 float32 = (3.14159 * 3.14159)
    v8 += v8
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
