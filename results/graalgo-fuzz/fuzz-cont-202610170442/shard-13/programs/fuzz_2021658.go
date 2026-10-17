package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64; F1 uint16; F2 int32 }
func genFn0(p0 bool) uint16 {
    var v6 string = "\n\t"
    var v7 int = 0
    v8 := GenS1{F0: 255}
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return 2
}
func genFn1() int32 {
    var v9 float64 = (-2.25 + (1e10 / 3.14159))
    v9 = 0.1
    _ = v9
    return -677
}

func main() {
    v1 := []string{"a", "foo bar", ""}
    var v2 int64 = 9223372036854775807
    v3 := make(map[int]int)
    var v4 int64 = (v2 + 42)
    var v5 int16 = -712
    v10 := GenS2{F0: 9223372036854775807, F1: 317, F2: 157}
    var v11 bool = (v4 > 970)
    for fi12 := 0; fi12 < 12; fi12++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        fmt.Println(v4)
        fmt.Println(fi12, v2, v2)
        _ = fi12
    }
    v13 := v3[128]
    fmt.Println(v5, v5)
    var v14 uint8 = 128
    var v15 float64 = float64(-2.25)
    var v16 float32 = (0.1 * (0.0 * 1e10))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
