package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
func genFn0(p0 int16, p1 int32, p2 uint8) bool {
    p0 = -1
    for fi4 := 0; fi4 < 3; fi4++ {
        var v5 int64 = 0
        var v6 float64 = 0.0
        _ = fi4
        _ = v5
        _ = v6
    }
    var v7 bool = (32767 != 2)
    var v8 string = "foo bar"
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return (-124 == 95)
}
func genFn1() int {
    if (0 <= 886) {
        if (128 <= 2) {
            v9 := []int{994, -829}
            _ = v9
        }
    } else {
        for fi10 := 0; fi10 < 15; fi10++ {
            for fi11 := 0; fi11 < 8; fi11++ {
                v12 := []int{128, fi10, 38, fi10}
                _ = fi11
                _ = v12
            }
            _ = fi10
        }
    }
    v13 := GenS1{F0: -1, F1: 423}
    _ = v13
    return 9223372036854775807
}

func main() {
    v1 := make(map[string]string)
    var v2 uint64 = 2
    v3 := make(map[string]bool)
    var v14 int32 = -573
    var v15 float32 = -2.25
    v16 := GenS1{F0: -98, F1: 291}
    var v17 float64 = (0.0 / 1.0)
    if true {
        var v18 int64 = -1
        var v19 uint64 = (v2 - (v2 | v2))
        _ = v18
        _ = v19
    }
    var v20 uint32 = 2147483647
    v3[("hello" + "hello")] = ((v20 ^ (v20 / (((v20 ^ (v20 >> ((2) & 63)))) | 1))) != v20)
    fmt.Println(v14)
    v20 -= 1
    fmt.Println(v14)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v3
}
