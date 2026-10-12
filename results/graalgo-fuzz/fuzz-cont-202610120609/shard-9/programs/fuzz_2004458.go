package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
type GenS2 struct { F0 int16 }

func main() {
    v1 := make(map[int]string)
    v2 := []string{}
    v3 := make(map[uint32]int)
    var v4 int32 = 0
    var v5 uint8 = 127
    for fi6 := 0; fi6 < 7; fi6++ {
        var v7 int = fi6
        var v8 bool = true
        switch fi6 {
        case 32767:
            for fi9 := 0; fi9 < 15; fi9++ {
                v4 = (v4 - v4)
                v8 = ((3.14159 - ((0.0 - 1.5) * 0.0)) == (0.1 * 1.5))
                _ = fi9
            }
        case -832:
            v7 = fi6
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    fmt.Println(v4, v5)
    v10 := GenS1{F0: (v5 | 2), F1: "foo bar"}
    _ = len(v3)
    v1[127] = (("foo bar" + "\n\t") + ("foo bar" + "hello"))
    v5 = (v5 >> ((830) & 63))
    switch 748 {
    case 127:
        v11 := GenS2{F0: -544}
        _ = v11
    case 32767:
        var v12 int16 = 32767
        _ = v12
    case 2147483647:
        if len(v2) > 0 {
            v2[0] = "a"
        }
    }
    _ = len(v2)
    v2 = append(v2, ("hello" + "\n\t"))
    if len(v2) > 0 {
        v2[0] = "\n\t"
    }
    if false {
        var v13 int32 = (v4 / ((v4) | 1))
        v5 = v5
        _ = v13
    } else {
        v14 := GenS1{F0: v5, F1: ("hello" + "\n\t")}
        _ = v14
    }
    _ = len(v2)
    if (0 != -39) {
        v1[2147483647] = "\n\t"
        v4 = v4
    }
    fmt.Println(v4, v4)
    var v15 int64 = 329
    if false {
        var v16 string = "foo bar"
        fmt.Println(v4)
        v5 = (v5 - v5)
        _ = v16
    } else {
        v17 := GenS1{F0: 32, F1: "\n\t"}
        _ = v17
    }
    v4 = (v4 * (v4 % ((v4) | 1)))
    var v18 float32 = 3.14159
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
