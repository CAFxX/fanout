package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 uint32 }

func main() {
    v1 := make(map[int]int)
    v2 := []string{}
    v3 := []string{"hello", "hello", ("foo bar" + ("a" + "foo bar"))}
    var v4 int = 1
    var v5 uint8 = 254
    v6 := []uint32{}
    _ = len(v2)
    var v7 int64 = 0
    var v8 int = v4
    var v9 uint32 = 4294967295
    v8 -= v4
    v6 = append(v6, v9)
    var v10 int64 = 42
    for fi11 := 0; fi11 < 8; fi11++ {
        v8 = (v4 & -279)
        fmt.Println(v10, v8, v4)
        _ = fi11
    }
    v12 := GenS1{F0: true, F1: "a", F2: 4294967295}
    switch v8 {
    case 42:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case 1:
        var v13 int64 = (^(v10))
        _ = v13
    }
    var v14 int = -610
    if len(v6) > 0 {
        v6[0] = 65535
    }
    v15 := GenS1{F0: true, F1: "foo bar", F2: v9}
    v6 = append(v6, v9)
    var v16 float32 = -2.25
    var v17 int32 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
