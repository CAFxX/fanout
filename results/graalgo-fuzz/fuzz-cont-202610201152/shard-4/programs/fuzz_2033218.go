package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 uint }
type GenS2 struct { F0 string; F1 bool }
func genFn0(p0 uint32, p1 bool, p2 int64) bool {
    var v4 float32 = 0.1
    v5 := GenS1{F0: p1, F1: 1, F2: 575}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return (42 <= 830)
}

func main() {
    v1 := []int{127, 9223372036854775807, 32767, 9223372036854775807}
    var v2 int16 = 963
    v3 := make(map[int]bool)
    for fi6 := 0; fi6 < 7; fi6++ {
        var v7 string = "hello"
        _ = fi6
        _ = v7
    }
    var v8 int64 = -9223372036854775808
    switch "a" {
    case "":
        _ = len(v3)
    default:
        v2 = (v2 - v2)
    }
    v9 := GenS2{F0: "\n\t", F1: ((32767 <= 2147483647) || (!false))}
    v10 := GenS2{F0: ("\n\t" + ""), F1: (0 >= v2)}
    var v11 int16 = (v2 / (((v2 ^ (v2 ^ int16(v8)))) | 1))
    v12 := GenS1{F0: true, F1: 255, F2: 561}
    v11 -= (^(v11))
    var v13 int32 = -1
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v8 = v8
    v13 -= v13
    v13 = v13
    v2 -= v2
    if len(v1) > 0 {
        v1[0] = 32767
    }
    var v14 int16 = -1
    v14 = v2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
