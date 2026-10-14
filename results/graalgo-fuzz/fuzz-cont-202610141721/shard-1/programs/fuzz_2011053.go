package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0(p0 int8) uint16 {
    var v5 string = "\n\t"
    v6 := GenS1{F0: 42}
    _ = p0
    _ = v5
    _ = v6
    return 1
}

func main() {
    var v1 float64 = 1.5
    v2 := make(map[int]bool)
    v3 := []int{}
    var v4 int64 = -9223372036854775808
    v1 += float64(0.0)
    if (v1 <= v1) {
        v7 := GenS1{F0: 338}
        fmt.Println(v1, v4)
        _ = v7
    } else {
        for fi8 := 0; fi8 < 1; fi8++ {
            var v9 uint16 = 2
            _ = fi8
            _ = v9
        }
    }
    for fi10 := 0; fi10 < 8; fi10++ {
        v2[-389] = ((true && false) || ((18446744073709551615 == 42) || true))
        var v11 string = "foo bar"
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 3; fi12++ {
        fmt.Println(v4, v4)
        _ = fi12
    }
    v4 = (v4 % ((0) | 1))
    fmt.Println(v4, v4)
    _ = len(v2)
    var v13 uint32 = 2147483648
    switch "\n\t" {
    case "":
        _ = len(v2)
    default:
        fmt.Println(v4, v13, v1)
    }
    v13 -= 1
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
