package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 uint64 }
func genFn0() uint8 {
    var v6 uint16 = 1
    switch "hello" {
    case "a":
        v6 += (v6 & v6)
    case "":
        v6 -= (v6 ^ (v6 % ((v6) | 1)))
    case "foo bar":
        var v7 int8 = 1
        _ = v7
    }
    v6 -= 957
    _ = v6
    return 238
}

func main() {
    var v1 float32 = 0.0
    var v2 int64 = -839
    var v3 uint = 0
    var v4 float32 = float32(0.0)
    var v5 uint16 = 620
    var v8 int64 = 1
    var v9 float32 = ((v1 + -2.25) / (1e10 / v1))
    if true {
        fmt.Println(v8, v1, v5)
    }
    v10 := []int{-2147483648, 127, -2147483648}
    v11 := GenS1{F0: 227, F1: -32768, F2: 127}
    v10 = append(v10, 4294967295)
    v8 += -1
    var v12 string = (((("a" + "\n\t") + "a") + "") + (("hello" + "\n\t") + "foo bar"))
    var v13 uint8 = 23
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
