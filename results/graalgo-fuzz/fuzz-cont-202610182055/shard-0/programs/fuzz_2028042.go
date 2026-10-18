package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    v1 := []string{"foo bar", "", ("" + "\n\t")}
    var v2 int32 = -2147483648
    var v3 uint64 = 704
    v4 := make(map[string]int)
    v5 := []int{-9223372036854775808, 4294967295, -1, 9223372036854775807}
    var v6 float32 = (0.1 - 0.1)
    var v7 uint = 1
    v3 += (v3 % (((^(v3))) | 1))
    if len(v1) > 0 {
        v1[0] = "foo bar"
    }
    var v8 bool = (!true)
    v6 += v6
    v9 := GenS1{F0: 417}
    fmt.Println(v6)
    var v10 uint32 = 762
    if true {
        v8 = (955 >= (v3 / ((v3) | 1)))
    }
    fmt.Println(v2, v6, v2)
    var v11 float32 = (1e10 / 1.0)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
