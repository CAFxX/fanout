package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }

func main() {
    var v1 uint64 = 430
    var v2 int16 = -552
    var v3 int16 = v2
    v4 := []uint32{42}
    var v5 bool = true
    var v6 uint = 653
    var v7 float32 = 1.5
    if len(v4) > 0 {
        v4[0] = 4294967295
    }
    if v5 {
        if len(v4) > 0 {
            v4[0] = 828
        }
        if len(v4) > 0 {
            v4[0] = 182
        }
    }
    var v8 uint8 = 42
    var v9 uint64 = (v1 - (v1 + (v1 % (((v1 | v1)) | 1))))
    switch "hello" {
    case "\n\t":
        v10 := GenS1{F0: -13, F1: 0}
        _ = v10
    case "a":
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    case "hello":
        var v11 uint32 = 2147483647
        _ = v11
    }
    var v12 uint64 = v9
    var v13 string = "foo bar"
    var v14 uint8 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
