package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 int }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int8; F1 int32 }

func main() {
    v1 := make(map[int]uint32)
    var v2 int = 653
    var v3 uint32 = 229
    var v4 int64 = -1
    var v5 uint16 = 412
    v6 := []int{}
    v7 := v1[32767]
    if false {
        var v8 float32 = 0.0
        v9 := GenS1{F0: 127, F1: ("foo bar" >= ("a" + ("" + ((("hello" + "") + ("hello" + "hello")) + "a")))), F2: v2}
        v2 -= -2147483648
        _ = v8
        _ = v9
    }
    var v10 int8 = 1
    v11 := GenS1{F0: v10, F1: false, F2: v2}
    v12 := GenS1{F0: -58, F1: false, F2: (-(v2))}
    if (true && ((v5 > (v5 % (((v5 & (v5 & 255))) | 1))) || (!false))) {
        switch "\n\t" {
        case "hello":
            v13 := GenS2{F0: v4}
            _ = v13
        case "":
            fmt.Println(v7, v7, v7)
        }
        fmt.Println(v2, v7, v2)
        fmt.Println(v5)
    }
    fmt.Println(v10, v4)
    if true {
        v14 := GenS3{F0: v10, F1: 993}
        _ = v14
    }
    v2 += 128
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v15 float32 = (0.1 - 0.1)
    v3 = 2147483647
    var v16 int16 = 596
    v1[(v2 - -630)] = (v7 << ((255) & 63))
    var v17 int = (v2 & v2)
    fmt.Println(v5, v3, v2)
    var v18 uint16 = (^(v5))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
