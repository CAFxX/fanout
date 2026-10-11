package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0() uint8 {
    var v4 int = -903
    var v5 uint8 = 0
    fmt.Println(v4, v5)
    _ = v4
    _ = v5
    return v5
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[string]uint32)
    v3 := make(map[uint32]bool)
    var v6 uint = 4294967295
    v6 = v6
    v6 -= (v6 / (((v6 >> ((v6) & 63))) | 1))
    for fi7 := 0; fi7 < 13; fi7++ {
        for fi8 := 0; fi8 < 15; fi8++ {
            v6 -= v6
            var v9 uint = (v6 + 255)
            _ = fi8
            _ = v9
        }
        v6 = (v6 % ((v6) | 1))
        v10 := GenS1{F0: -1}
        _ = fi7
        _ = v10
    }
    v6 = v6
    var v11 uint32 = 65535
    fmt.Println(v6, v11)
    if false {
        delete(v3, v11)
        v3[(v11 / (((v11 ^ (v11 / ((v11) | 1)))) | 1))] = true
        var v12 int = -9223372036854775808
        _ = v12
    } else {
        for fi13 := 0; fi13 < 12; fi13++ {
            v14 := []int{(fi13 - (fi13 + -2147483648)), fi13, fi13}
            _ = fi13
            _ = v14
        }
        v15 := []int{}
        _ = v15
    }
    if ((true || true) || false) {
        switch 1 {
        case 32768:
            v6 += v6
        default:
            fmt.Println(v6, v6)
        }
    }
    v2["foo bar"] = (v11 * 4294967295)
    v11 += v11
    v6 += v6
    v16 := []int{1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v16
    _ = v2
    _ = v3
    _ = v6
}
