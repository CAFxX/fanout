package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 bool; F1 int64 }
func genFn0(p0 int16) int16 {
    v5 := []int{42, -2147483648, 42, 1}
    v6 := GenS1{F0: "", F1: 243}
    p0 += (p0 - p0)
    var v7 uint64 = 2
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return (p0 >> ((744) & 63))
}

func main() {
    v1 := []int{2147483647, -294, 127, 315}
    var v2 bool = true
    var v3 uint8 = 42
    v4 := make(map[uint32]int)
    var v8 uint8 = (v3 << ((1) & 63))
    fmt.Println(v3)
    if ("foo bar" <= ("a" + "foo bar")) {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v8 = v3
        if v2 {
            var v9 float64 = -2.25
            for fi10 := 0; fi10 < 20; fi10++ {
                v1 = append(v1, 857)
                v11 := GenS1{F0: "a", F1: v3}
                fmt.Println(v9, v2, v2)
                _ = fi10
                _ = v11
            }
            fmt.Println(v8, v2, v2)
            _ = v9
        }
    }
    _ = len(v4)
    if v2 {
        for fi12 := 0; fi12 < 13; fi12++ {
            for fi13 := 0; fi13 < 10; fi13++ {
                v3 = (v8 & (v8 + (^(v8))))
                _ = fi13
            }
            fmt.Println(v2, v2)
            _ = fi12
        }
        if true {
            v3 = (v8 % ((v8) | 1))
        }
        switch 892 {
        case -2147483648:
            fmt.Println(v2, v3, v3)
        case 127:
            v8 += 95
        default:
            v14 := GenS2{F0: (65535 < 972), F1: 42}
            _ = v14
        }
    }
    var v15 uint16 = 335
    fmt.Println(v8)
    v8 = (v8 << ((255) & 63))
    v16 := GenS2{F0: v2, F1: 0}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
