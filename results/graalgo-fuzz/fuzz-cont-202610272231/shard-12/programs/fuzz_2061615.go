package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    v1 := []int{127, 0, -9223372036854775808}
    var v2 string = (("hello" + "") + (("foo bar" + "hello") + "\n\t"))
    var v3 uint16 = 42
    v4 := make(map[int]uint32)
    var v5 int16 = 1000
    v6 := make(map[string]int)
    var v7 float64 = (((float64((1e10 + (0.1 * 1.5))) - -2.25) * (0.1 + float64(0.0))) + 1.5)
    var v8 int = -1
    v7 = float64(float32(v7))
    v9 := GenS1{F0: 1}
    v10 := GenS1{F0: -2147483648}
    if len(v1) > 0 {
        v1[0] = (v8 / (((v8 | (v8 << ((401) & 63)))) | 1))
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
