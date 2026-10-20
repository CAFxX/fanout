package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
func genFn0(p0 int) int32 {
    fmt.Println(p0)
    v5 := []int{(p0 - (p0 / ((p0) | 1)))}
    _ = p0
    _ = v5
    return 42
}

func main() {
    var v1 uint8 = 127
    var v2 int32 = 326
    var v3 string = "hello"
    v4 := make(map[int]bool)
    if (32768 > -71) {
        v6 := GenS1{F0: "foo bar", F1: v2}
        fmt.Println(v1)
        _ = v6
    }
    if false {
        v7 := GenS1{F0: v3, F1: (v2 + (v2 >> ((18446744073709551615) & 63)))}
        if (0.0 == float64(float32(1.5))) {
            v3 = (v3 + v3)
            _ = len(v4)
        }
        _ = v7
    } else {
        var v8 bool = (false && (true && false))
        v9 := GenS1{F0: v3, F1: -2147483648}
        var v10 int32 = (v2 - genFn0(2147483647))
        _ = v10
        _ = v8
        _ = v9
    }
    v1 += 42
    if (42 > 42) {
        v2 = 182
        fmt.Println(v1, v1)
    }
    if false {
        v11 := GenS1{F0: v3, F1: v2}
        var v12 float32 = 1e10
        _ = v11
        _ = v12
    }
    v2 = (v2 % ((genFn0(-875)) | 1))
    v13 := []int{4294967295, 127}
    var v14 bool = (true || true)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
