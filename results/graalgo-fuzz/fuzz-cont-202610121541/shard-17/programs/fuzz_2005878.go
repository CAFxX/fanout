package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 uint32 }
type GenS2 struct { F0 uint16; F1 int8 }

func main() {
    var v1 uint = 456
    v2 := make(map[uint32]uint32)
    v3 := []string{("hello" + ("\n\t" + "hello"))}
    var v4 uint = (v1 << ((255) & 63))
    v5 := GenS1{F0: -696, F1: ("a" + ("a" + (("" + ("" + ("\n\t" + ""))) + ("" + ("\n\t" + ("a" + "")))))), F2: 109}
    fmt.Println(v4, v1)
    switch 32768 {
    case 0:
        var v6 float32 = 0.1
        _ = v6
    case 128:
        v4 = v4
    default:
        var v7 int64 = -1
        _ = v7
    }
    for fi8 := 0; fi8 < 14; fi8++ {
        v4 += (v4 | 255)
        v9 := GenS2{F0: 2, F1: 0}
        _ = fi8
        _ = v9
    }
    v1 -= v1
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v10 := GenS2{F0: 450, F1: -1}
    v11 := v2[65535]
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
