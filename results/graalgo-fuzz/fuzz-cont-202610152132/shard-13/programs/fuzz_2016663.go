package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 int }

func main() {
    var v1 bool = ((((float32(1e10) - 0.0) - -2.25) + (0.1 / 1.0)) <= 1e10)
    v2 := make(map[int]string)
    var v3 int8 = 42
    var v4 int16 = 192
    v5 := []int{-642}
    v2[-9223372036854775808] = (("hello" + ("a" + "")) + ("a" + "hello"))
    fmt.Println(v3, v4, v4)
    var v6 int8 = v3
    for fi7 := 0; fi7 < 3; fi7++ {
        for fi8 := 0; fi8 < 18; fi8++ {
            fmt.Println(fi7, v6, fi8)
            v3 += (v3 - (v3 & v6))
            _ = fi8
        }
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = fi7
    }
    v9 := GenS1{F0: 43, F1: "", F2: -1}
    v10 := v2[-217]
    v3 = v6
    delete(v2, 84)
    v6 = v6
    fmt.Println(v10, v10, v6)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
