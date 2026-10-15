package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
func genFn0(p0 uint32, p1 uint16) int32 {
    switch 32767 {
    case 4294967295:
        fmt.Println(p1, p0)
    default:
        var v7 int = -9223372036854775808
        _ = v7
    }
    fmt.Println(p0)
    p1 = (p1 ^ p1)
    switch "hello" {
    case "a":
        p0 = p0
    default:
        v8 := []int{0, 4294967295, 32767, 720}
        _ = v8
    }
    v9 := []int{}
    _ = p0
    _ = p1
    _ = v9
    return 223
}

func main() {
    var v1 int16 = 187
    var v2 int8 = 42
    var v3 uint16 = 32767
    v4 := make(map[int]int)
    var v5 bool = (128 > 32767)
    var v6 float64 = 0.1
    v6 += v6
    v10 := []int{-1, -611}
    if (1 <= 526) {
        switch "\n\t" {
        case "a":
            v5 = v5
        case "\n\t":
            var v11 int16 = 32767
            _ = v11
        }
        switch ("hello" + ("hello" + "foo bar")) {
        case "hello":
            _ = len(v10)
        default:
            if (v1 != v1) {
                v6 = v6
            }
        }
        v12 := GenS1{F0: 1000, F1: 65535}
        _ = v12
    }
    fmt.Println(v1, v5, v2)
    if (-1 >= 988) {
        v13 := GenS1{F0: 152, F1: (v3 + (v3 ^ v3))}
        _ = v13
    }
    fmt.Println(v2)
    v14 := GenS1{F0: v1, F1: 255}
    var v15 int16 = 1000
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
