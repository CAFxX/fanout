package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }

func main() {
    v1 := make(map[uint32]string)
    var v2 int = 2147483647
    var v3 float64 = 3.14159
    v4 := make(map[uint32]string)
    var v5 uint = 42
    var v6 string = ("" + "")
    var v7 string = v6
    v8 := []int{4294967295, v2}
    var v9 int64 = -569
    var v10 bool = false
    fmt.Println(v6)
    fmt.Println(v6, v10, v7)
    v10 = (1e10 == ((0.0 + (0.1 + 3.14159)) - ((1e10 * -2.25) * 0.1)))
    fmt.Println(v3, v7, v2)
    v11 := GenS1{F0: (v2 | (v2 % (((-(v2))) | 1))), F1: -1}
    if (!((!(32767 != -32768)) || (-15 >= 12))) {
        v12 := GenS1{F0: 156, F1: 0}
        fmt.Println(v7, v10)
        fmt.Println(v10)
        _ = v12
    }
    if (v2 <= (v2 * 128)) {
        v9 = int64(v5)
    } else {
        var v13 int = 0
        delete(v4, 2)
        _ = v13
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
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
