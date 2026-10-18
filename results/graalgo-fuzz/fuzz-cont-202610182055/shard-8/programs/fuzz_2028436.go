package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 int }
func genFn0(p0 uint16, p1 int32) uint64 {
    v4 := []int{}
    p1 = (p1 | p1)
    p1 -= -646
    if true {
        p0 -= p0
    }
    p1 -= p1
    _ = p0
    _ = p1
    _ = v4
    return 2
}

func main() {
    v1 := []uint32{42, 928, 0, 2}
    var v2 int16 = 648
    v3 := make(map[uint32]int)
    v5 := v3[2]
    v5 = v5
    v5 = v5
    var v6 bool = false
    var v7 int = (v5 % ((v5) | 1))
    var v8 uint64 = 135
    fmt.Println(v8, v2)
    fmt.Println(v8)
    v7 -= 716
    var v9 bool = false
    switch (v7 | 127) {
    case 32767:
        v10 := GenS1{F0: v8, F1: ("foo bar" + "\n\t"), F2: (^(v5))}
        _ = v10
    default:
        v6 = ((-(v7)) != (v7 * v7))
    }
    if false {
        var v11 int16 = (v2 % (((v2 | -47)) | 1))
        _ = v11
    }
    fmt.Println(v5, v6)
    v8 = (v8 * (v8 - 940))
    v5 = v5
    switch -191 {
    case -1:
        var v12 int64 = -1
        _ = v12
    case 127:
        fmt.Println(v9, v9, v9)
    case 32767:
        v9 = (!(3.14159 <= 3.14159))
    default:
        var v13 float32 = 1e10
        _ = v13
    }
    v6 = (false && (-1 > int8(v8)))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
