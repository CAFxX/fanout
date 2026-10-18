package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 int16 }
func genFn0(p0 uint8, p1 uint8) int {
    var v6 bool = true
    p1 -= p0
    var v7 uint8 = p1
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return -40
}

func main() {
    v1 := []string{"", "a", "hello"}
    v2 := []int{0, -619}
    var v3 float64 = (((((1e10 / 1e10) + ((0.0 + (-2.25 - float64(-2.25))) / (0.1 + 3.14159))) / -2.25) / 0.1) - (1.5 - (3.14159 / 0.1)))
    v4 := make(map[int]bool)
    v5 := []string{"foo bar"}
    var v8 uint64 = 1
    if true {
        switch genFn0(1, 255) {
        case 9223372036854775807:
            v9 := GenS1{F0: 4294967295, F1: "a", F2: 1000}
            _ = v9
        case -589:
            var v10 float32 = 0.0
            _ = v10
        case 32767:
            for fi11 := 0; fi11 < 10; fi11++ {
                _ = len(v4)
                _ = fi11
            }
        }
        var v12 uint = 65535
        _ = v12
    } else {
        v4[-626] = (255 != 44)
        v8 = (v8 << ((649) & 63))
    }
    v13 := GenS1{F0: 42, F1: "a", F2: -70}
    var v14 int64 = 1
    v14 = (v14 + 1)
    v15 := GenS1{F0: 32767, F1: ("" + ("a" + ("foo bar" + "\n\t"))), F2: -1}
    if len(v2) > 0 {
        v2[0] = 1
    }
    switch 32767 {
    case 2147483647:
        fmt.Println(v14)
    default:
        fmt.Println(v14, v3, v8)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
