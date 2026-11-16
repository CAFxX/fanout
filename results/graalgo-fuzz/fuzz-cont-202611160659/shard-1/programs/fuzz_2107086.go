package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 int64 }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[uint32]int)
    var v3 uint32 = 1
    var v4 int8 = 1
    delete(v2, 290)
    v4 = (^(v4))
    if (65535 <= (v3 & v3)) {
        v3 = (v3 & 821)
        var v5 int8 = (v4 >> ((584) & 63))
        _ = v5
    }
    var v6 uint32 = v3
    var v7 uint16 = 32767
    v8 := []int{1, 4294967295, 2147483647}
    delete(v2, (v6 >> ((42) & 63)))
    var v9 int8 = 127
    var v10 float32 = 0.0
    if true {
        v10 = v10
        if true {
            var v11 uint32 = (v3 % (((^(v6))) | 1))
            fmt.Println(v10, v4)
            fmt.Println(v7, v7)
            _ = v11
        } else {
            v10 -= (((0.0 + (float32(0.0) + v10)) / v10) + float32(3.14159))
        }
        v12 := v2[(v3 - v6)]
        _ = v12
    }
    v13 := v1["a"]
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    v14 := GenS1{F0: -83, F1: ("foo bar" + "\n\t"), F2: -1}
    fmt.Println(v6, v3)
    delete(v1, "a")
    var v15 int = (^(v13))
    var v16 float32 = ((v10 - v10) - ((((3.14159 * v10) + 0.1) + 0.1) / (((v10 * 3.14159) + v10) - 1.5)))
    v13 -= v13
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
