package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }

func main() {
    v1 := []string{"", "foo bar"}
    var v2 int32 = 2147483647
    var v3 int8 = 30
    if ((float32(0.0) * -2.25) < 0.1) {
        fmt.Println(v2, v2)
    }
    v2 = int32(v3)
    v2 = (v2 / ((1) | 1))
    switch "\n\t" {
    case "hello":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case "foo bar":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v4 uint8 = 1
    v3 += v3
    if false {
        fmt.Println(v2, v4, v2)
        if true {
            v3 = v3
        }
        _ = len(v1)
    }
    v5 := GenS1{F0: 2147483647, F1: 717}
    var v6 uint = 1
    v4 -= v4
    v2 = (v2 % (((v2 & (v2 & 1))) | 1))
    v7 := GenS1{F0: 1, F1: 43}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
