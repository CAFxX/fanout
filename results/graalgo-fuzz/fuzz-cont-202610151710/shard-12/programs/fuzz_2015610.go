package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 uint64 }
type GenS2 struct { F0 uint32 }

func main() {
    v1 := make(map[int]int)
    var v2 bool = (42 != -923)
    var v3 uint64 = 815
    v4 := []string{"foo bar", "a", "hello", ("hello" + "")}
    var v5 bool = true
    v6 := GenS1{F0: -1, F1: -601, F2: v3}
    v4 = append(v4, "")
    var v7 int64 = -934
    v3 -= v3
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v8 int8 = -1
    var v9 uint8 = 127
    v4 = append(v4, "hello")
    var v10 uint32 = 855
    fmt.Println(v2, v3, v2)
    if len(v4) > 0 {
        v4[0] = "a"
    }
    var v11 int32 = -2147483648
    if false {
        switch "a" {
        case "":
            v12 := GenS1{F0: 42, F1: 421, F2: v3}
            _ = v12
        default:
            v13 := GenS2{F0: (v10 ^ (v10 ^ (v10 & 2)))}
            _ = v13
        }
        var v14 float32 = 0.0
        _ = v14
    } else {
        v11 = v11
    }
    v4 = append(v4, "\n\t")
    var v15 float32 = -2.25
    v7 += (v7 % ((42) | 1))
    var v16 uint = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
