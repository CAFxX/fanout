package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
func genFn0(p0 int8, p1 bool) int {
    v4 := GenS1{F0: (1 >= -9223372036854775808), F1: -4}
    var v5 uint16 = 19
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 9223372036854775807
}

func main() {
    var v1 int32 = -1
    var v2 string = "hello"
    var v3 float64 = (0.1 * 1e10)
    for fi6 := 0; fi6 < 16; fi6++ {
        var v7 uint16 = 255
        _ = fi6
        _ = v7
    }
    if (((v3 + v3) * 0.1) <= v3) {
        v2 = v2
        v2 = v2
    } else {
        fmt.Println(v2, v1, v3)
    }
    var v8 int16 = 14
    v9 := GenS1{F0: true, F1: 127}
    v1 -= 1
    for fi10 := 0; fi10 < 20; fi10++ {
        fmt.Println(fi10)
        _ = fi10
    }
    v2 = v2
    switch v2 {
    case "foo bar":
        v11 := []int{}
        _ = v11
    case "":
        v1 = -314
    default:
        v12 := []int{-281, 32768, 32767}
        _ = v12
    }
    v13 := GenS1{F0: (v8 == v8), F1: 57}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
