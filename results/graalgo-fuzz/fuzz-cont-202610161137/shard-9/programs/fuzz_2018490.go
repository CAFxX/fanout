package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }

func main() {
    v1 := make(map[int]uint32)
    var v2 string = "a"
    v3 := []string{}
    var v4 uint = 0
    var v5 uint = 325
    v6 := []int{-693}
    var v7 int16 = -1
    fmt.Println(v2, v2)
    v5 = v4
    _ = len(v1)
    for fi8 := 0; fi8 < 8; fi8++ {
        v9 := GenS1{F0: (("a" + v2) + v2), F1: (v7 | (v7 % (((v7 & (v7 - (v7 >> ((v4) & 63))))) | 1)))}
        var v10 bool = false
        v7 += (v7 + 1)
        _ = fi8
        _ = v10
        _ = v9
    }
    var v11 uint8 = 240
    switch -403 {
    case -242:
        v7 -= (v7 & (-(v7)))
    case 42:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    default:
        var v12 float64 = -2.25
        _ = v12
    }
    switch 4294967295 {
    case 983:
        var v13 string = v2
        _ = v13
    case 460:
        v2 = (v2 + ("\n\t" + v2))
    default:
        v5 = (v4 & 65535)
    }
    var v14 int64 = 570
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
