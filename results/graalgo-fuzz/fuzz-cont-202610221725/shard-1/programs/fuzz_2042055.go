package main

import (
    "fmt"
)


func main() {
    var v1 int64 = -666
    v2 := []int{127, 32767}
    var v3 uint = 96
    v4 := []int{-2147483648, -2147483648, 35}
    v5 := make(map[int]string)
    var v6 uint8 = 255
    v3 += (v3 ^ v3)
    v6 = 255
    v3 = (v3 % ((v3) | 1))
    v3 -= (v3 ^ v3)
    switch "\n\t" {
    case "hello":
        fmt.Println(v6, v6, v3)
    case "\n\t":
        fmt.Println(v3)
    }
    var v7 int64 = -236
    v6 = v6
    v5[4294967295] = ("\n\t" + "a")
    var v8 uint64 = 18446744073709551615
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
