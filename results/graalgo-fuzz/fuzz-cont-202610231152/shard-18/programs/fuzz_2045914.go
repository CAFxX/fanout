package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 string }
type GenS2 struct { F0 int16; F1 string }
func genFn0(p0 string, p1 uint64) int {
    v4 := []int{32767, 0, 645, 42}
    var v5 int64 = -9223372036854775808
    for fi6 := 0; fi6 < 3; fi6++ {
        v7 := GenS1{F0: -2147483648, F1: 626, F2: "foo bar"}
        _ = len(v4)
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 2147483647
}

func main() {
    v1 := make(map[int]bool)
    v2 := []string{}
    var v3 float64 = 0.0
    for fi8 := 0; fi8 < 3; fi8++ {
        if false {
            fmt.Println(fi8, v3)
        } else {
            var v9 uint = 528
            v10 := v1[0]
            _ = v10
            _ = v9
        }
        _ = fi8
    }
    var v11 uint8 = 1
    var v12 int16 = -773
    var v13 int = genFn0(("\n\t" + ""), 2)
    fmt.Println(v13, v13)
    var v14 uint16 = 603
    fmt.Println(v12, v11, v12)
    if true {
        v13 = v13
        var v15 uint16 = (^(v14))
        _ = v15
    } else {
        v16 := v1[v13]
        _ = v16
    }
    var v17 int32 = int32(v12)
    v11 += (v11 * v11)
    fmt.Println(v11, v13)
    fmt.Println(v14, v12)
    if true {
        fmt.Println(v14)
    } else {
        for fi18 := 0; fi18 < 10; fi18++ {
            v19 := GenS2{F0: int16(v11), F1: ("foo bar" + "hello")}
            v12 += (v12 << ((2) & 63))
            var v20 bool = ((((65535 >= 1) || true) || (true && false)) && ((v11 + 127) < (v11 >> ((2) & 63))))
            _ = fi18
            _ = v19
            _ = v20
        }
    }
    if true {
        delete(v1, (v13 ^ (v13 - 935)))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v17
    _ = v2
    _ = v3
}
