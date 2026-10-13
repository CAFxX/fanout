package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 uint32 }
type GenS2 struct { F0 uint32; F1 int }

func main() {
    v1 := make(map[string]int)
    var v2 uint32 = 2
    var v3 int8 = 0
    var v4 string = "hello"
    var v5 bool = (v2 == 57)
    var v6 uint32 = (v2 >> ((569) & 63))
    v6 = 42
    switch 2147483647 {
    case 39:
        fmt.Println(v4, v5)
    case 32768:
        v7 := GenS1{F0: 4294967295, F1: -495, F2: 2147483647}
        _ = v7
    case 52:
        switch -9223372036854775808 {
        case -891:
            v6 += (v6 * (^(v2)))
        }
    default:
        v8 := GenS2{F0: 19, F1: -2147483648}
        _ = v8
    }
    if (!(("a" + "hello") > (("" + v4) + (v4 + "a")))) {
        v9 := GenS1{F0: 1, F1: 9223372036854775807, F2: v2}
        _ = len(v1)
        var v10 float32 = 3.14159
        _ = v10
        _ = v9
    }
    if (794 > -117) {
        fmt.Println(v6, v2, v3)
        var v11 bool = v5
        v3 = (v3 >> ((229) & 63))
        _ = v11
    }
    v12 := []int{1}
    for fi13 := 0; fi13 < 15; fi13++ {
        if v5 {
            if len(v12) > 0 {
                v12[0] = (fi13 << ((255) & 63))
            }
        }
        _ = fi13
    }
    v5 = true
    var v14 uint64 = 18446744073709551615
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
