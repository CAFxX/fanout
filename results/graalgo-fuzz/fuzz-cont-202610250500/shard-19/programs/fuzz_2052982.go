package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint }

func main() {
    var v1 int = 2147483647
    v2 := make(map[int]int)
    v3 := []int{}
    v4 := make(map[uint32]int)
    var v5 bool = false
    v6 := []string{"foo bar"}
    if (0 >= 65535) {
        v1 -= 127
    } else {
        v7 := v2[(v1 ^ (v1 - v1))]
        fmt.Println(v1, v1)
        v3 = append(v3, v7)
        _ = v7
    }
    var v8 float64 = 1.5
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v1 = v1
    v9 := v2[v1]
    var v10 int64 = 245
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v5 = ((v10 >> ((612) & 63)) == v10)
    v11 := GenS1{F0: 59}
    v12 := GenS1{F0: 1}
    var v13 uint = 2
    var v14 uint = v13
    fmt.Println(v8)
    v15 := GenS2{F0: 375}
    delete(v2, v9)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
