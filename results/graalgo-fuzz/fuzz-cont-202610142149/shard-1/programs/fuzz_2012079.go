package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }

func main() {
    var v1 float32 = (1e10 / -2.25)
    var v2 uint8 = 19
    var v3 int8 = -128
    v4 := make(map[int]string)
    v5 := make(map[string]string)
    var v6 int8 = (v3 - (^(v3)))
    switch -1 {
    case 32768:
        if (float64(v1) <= 0.1) {
            _ = len(v4)
        }
    case -713:
        switch 42 {
        case 639:
            v2 -= (v2 % ((v2) | 1))
        }
    case 127:
        var v7 int8 = v3
        _ = v7
    }
    var v8 string = ("" + "a")
    fmt.Println(v3)
    v9 := GenS1{F0: v8, F1: -1}
    v10 := []int{4294967295, 32767, -2147483648, 347}
    delete(v4, 563)
    v8 = "hello"
    if (!true) {
        var v11 uint32 = 768
        var v12 uint32 = (v11 / ((0) | 1))
        if (((1.5 * (1e10 / (((0.0 + float64((v1 - v1))) * 1e10) - 3.14159))) * (0.0 * (0.0 / 1.5))) != 0.0) {
            v8 = ""
            v8 = "foo bar"
            var v13 int16 = int16(v6)
            _ = v13
        }
        _ = v11
        _ = v12
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
