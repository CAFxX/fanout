package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[uint32]int)
    var v3 uint16 = 380
    var v4 uint16 = v3
    var v5 int64 = 9223372036854775807
    var v6 string = "hello"
    v4 += v4
    var v7 uint = 1
    switch 903 {
    case 9223372036854775807:
        v4 = (v3 + (^(v3)))
    case -1:
        v8 := GenS1{F0: 96, F1: 32767}
        _ = v8
    }
    v7 = 556
    for fi9 := 0; fi9 < 12; fi9++ {
        v4 = (v3 / ((v4) | 1))
        var v10 int = fi9
        _ = fi9
        _ = v10
    }
    v5 += (^(v5))
    var v11 uint8 = 128
    for fi12 := 0; fi12 < 9; fi12++ {
        v4 -= v4
        for fi13 := 0; fi13 < 9; fi13++ {
            if true {
                fmt.Println(v6, v4, v11)
            }
            _ = fi13
        }
        _ = fi12
    }
    switch "hello" {
    case "\n\t":
        switch 814 {
        case 4294967295:
            v14 := []int{}
            _ = v14
        case -9223372036854775808:
            var v15 uint = v7
            _ = v15
        }
    default:
        if (-2147483648 >= -731) {
            v16 := GenS1{F0: -95, F1: -2147483648}
            _ = v16
        } else {
            v17 := []int{32768}
            _ = v17
        }
    }
    var v18 int64 = (-(v5))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
