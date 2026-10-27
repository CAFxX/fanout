package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
func genFn0() bool {
    var v6 uint32 = 65535
    fmt.Println(v6)
    switch 127 {
    case 327:
        fmt.Println(v6)
    case -1:
        var v7 uint32 = (v6 % ((v6) | 1))
        _ = v7
    case 512:
        v8 := []int{128, -841, 42, 127}
        _ = v8
    }
    for fi9 := 0; fi9 < 18; fi9++ {
        v6 -= 828
        var v10 int16 = 230
        v11 := []int{}
        _ = fi9
        _ = v10
        _ = v11
    }
    _ = v6
    return (!true)
}

func main() {
    var v1 uint = 1
    v2 := []string{"hello", "a", ("a" + "")}
    var v3 int16 = 32767
    var v4 uint32 = 1
    var v5 float32 = (1.5 * -2.25)
    v1 += (v1 | v1)
    if true {
        v2 = append(v2, "a")
        var v12 int32 = 0
        _ = v12
    } else {
        v3 += v3
    }
    v13 := GenS1{F0: uint64(v1), F1: 4294967295}
    v4 = v4
    fmt.Println(v4, v3, v3)
    if true {
        fmt.Println(v5, v4)
    }
    v1 -= 255
    var v14 int64 = 1
    v14 += v14
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println(v1, v14)
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
