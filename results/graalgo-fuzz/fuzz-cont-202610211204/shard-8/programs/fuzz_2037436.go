package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 int8 }

func main() {
    var v1 int64 = 943
    v2 := []int{}
    var v3 int8 = 67
    var v4 uint = 65535
    v5 := GenS1{F0: 1, F1: (v4 / ((v4) | 1)), F2: v3}
    v4 = 507
    var v6 string = "\n\t"
    var v7 bool = ((v1 / (((v1 * v1)) | 1)) > v1)
    var v8 uint8 = 169
    v3 = (v3 - (v3 * (v3 % ((-93) | 1))))
    v2 = append(v2, 128)
    fmt.Println(v4, v3, v3)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
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
