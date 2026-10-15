package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 uint }
type GenS2 struct { F0 bool; F1 uint32 }

func main() {
    var v1 uint32 = 65535
    var v2 bool = true
    v3 := []int{127, -1, -26, 42}
    var v4 bool = v2
    v2 = v2
    var v5 int8 = 42
    v5 -= (v5 - v5)
    var v6 bool = (2 > 421)
    switch ("foo bar" + ("hello" + "")) {
    case "foo bar":
        v7 := GenS1{F0: v1, F1: 143, F2: 2}
        _ = v7
    default:
        v5 += v5
    }
    _ = len(v3)
    var v8 bool = (513 != 42)
    _ = len(v3)
    v9 := GenS2{F0: (214 <= 641), F1: v1}
    v10 := GenS2{F0: (v2 && v8), F1: v1}
    fmt.Println(v4, v6)
    var v11 uint32 = (v1 * 2147483648)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
