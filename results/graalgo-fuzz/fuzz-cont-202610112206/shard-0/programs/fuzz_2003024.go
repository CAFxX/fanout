package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }

func main() {
    v1 := make(map[uint32]string)
    var v2 int32 = 0
    v3 := make(map[string]uint32)
    v4 := make(map[string]int)
    v5 := make(map[uint32]int)
    v6 := make(map[string]string)
    v2 = v2
    v7 := GenS1{F0: false, F1: true}
    v2 = (v2 % ((v2) | 1))
    if true {
        var v8 int = -1
        v2 += -1
        v9 := GenS1{F0: (!(!("a" == (("\n\t" + "a") + "")))), F1: false}
        _ = v8
        _ = v9
    } else {
        v10 := []int{32767, -2147483648, 2147483647, 127}
        if (!true) {
            fmt.Println(v2)
            v2 += (^(v2))
            if len(v10) > 0 {
                v10[0] = 9223372036854775807
            }
        }
        var v11 int = 2147483647
        _ = v10
        _ = v11
    }
    var v12 string = "hello"
    v2 -= v2
    var v13 float32 = -2.25
    v2 += (v2 - v2)
    v14 := []int{574, 592, 127, -980}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
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
}
