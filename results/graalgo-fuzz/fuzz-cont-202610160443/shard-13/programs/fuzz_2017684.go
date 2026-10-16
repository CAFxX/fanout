package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]bool)
    v3 := []int{}
    var v4 bool = true
    v5 := GenS1{F0: "", F1: 65535}
    fmt.Println(v4)
    var v6 int8 = -43
    var v7 int16 = -1
    var v8 uint16 = 2
    for fi9 := 0; fi9 < 8; fi9++ {
        if len(v3) > 0 {
            v3[0] = fi9
        }
        v4 = true
        var v10 bool = false
        _ = fi9
        _ = v10
    }
    v4 = (0 > 0)
    v4 = (float64(1.5) >= 3.14159)
    var v11 uint = 42
    var v12 float64 = 1.5
    var v13 uint = v11
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v7 -= (v7 ^ (v7 - -600))
    if (!false) {
        var v14 int32 = 1
        _ = v14
    } else {
        fmt.Println(v13, v11, v11)
        var v15 string = "foo bar"
        for fi16 := 0; fi16 < 10; fi16++ {
            v11 -= v11
            _ = fi16
        }
        _ = v15
    }
    var v17 bool = (false && (9223372036854775807 < -9223372036854775808))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
