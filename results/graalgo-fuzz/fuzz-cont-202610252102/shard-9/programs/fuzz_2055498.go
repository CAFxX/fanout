package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
func genFn0(p0 int8, p1 int16, p2 string) int64 {
    v7 := GenS1{F0: 2147483647, F1: "foo bar"}
    v8 := []int{0, 854}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 42
}

func main() {
    v1 := []string{"", "foo bar", "\n\t", "hello"}
    v2 := make(map[uint32]int)
    var v3 int = 132
    var v4 int32 = 1
    var v5 float64 = 0.0
    var v6 float64 = ((float64(-2.25) + float64(-2.25)) - v5)
    fmt.Println(v6)
    v9 := GenS1{F0: 961, F1: "hello"}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v10 uint = 255
    var v11 uint64 = 1
    v12 := GenS1{F0: 4294967295, F1: ""}
    var v13 bool = false
    var v14 float64 = v5
    v6 -= (((float64(((float32(-2.25) / 0.1) * 3.14159)) * 1.5) * (((v6 + v14) * v5) + (0.0 / -2.25))) * (v14 - (v5 + 1.5)))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
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
    _ = v9
}
