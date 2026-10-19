package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    v1 := []string{"hello"}
    var v2 int32 = 0
    var v3 uint8 = 127
    var v4 string = "\n\t"
    v5 := []string{"hello", "\n\t", v4}
    v6 := make(map[int]bool)
    v7 := GenS1{F0: ("foo bar" > ("a" + ("a" + v4)))}
    v4 = v4
    v6[-2147483648] = (42 > 916)
    fmt.Println(v2, v2)
    var v8 int32 = (v2 | (-(v2)))
    for fi9 := 0; fi9 < 12; fi9++ {
        fmt.Println(fi9, fi9, v8)
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = fi9
    }
    var v10 int64 = 672
    var v11 float64 = (3.14159 + 1.5)
    var v12 int8 = 0
    v2 += int32(v3)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v13 uint = 2
    v11 = v11
    v11 += -2.25
    v10 += (v10 & (v10 % ((v10) | 1)))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
