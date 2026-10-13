package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }
func genFn0() string {
    for fi4 := 0; fi4 < 20; fi4++ {
        fmt.Println(fi4)
        var v5 int8 = 29
        _ = fi4
        _ = v5
    }
    var v6 float64 = 0.1
    _ = v6
    return "foo bar"
}

func main() {
    var v1 bool = (43 > 42)
    v2 := make(map[string]string)
    var v3 uint32 = 0
    switch 128 {
    case -9223372036854775808:
        v1 = v1
    case -949:
        var v7 uint8 = 42
        _ = v7
    case 4294967295:
        fmt.Println(v1, v1)
    }
    v1 = true
    var v8 float64 = (float64(1.5) - 1e10)
    v9 := []int{}
    v10 := GenS1{F0: -895, F1: ("" + "foo bar")}
    var v11 int32 = 466
    for fi12 := 0; fi12 < 12; fi12++ {
        var v13 string = "\n\t"
        var v14 uint8 = 127
        if len(v9) > 0 {
            v9[0] = fi12
        }
        _ = fi12
        _ = v13
        _ = v14
    }
    var v15 string = "foo bar"
    var v16 int32 = v11
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
