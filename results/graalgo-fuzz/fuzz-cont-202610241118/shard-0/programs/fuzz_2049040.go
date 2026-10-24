package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
func genFn0() uint64 {
    if ((!true) && true) {
        var v5 uint32 = 921
        v5 = (^(v5))
        var v6 int = -984
        _ = v5
        _ = v6
    }
    for fi7 := 0; fi7 < 4; fi7++ {
        var v8 bool = false
        v8 = v8
        _ = fi7
        _ = v8
    }
    return 1
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint16 = 1
    var v3 int32 = 2147483647
    v4 := []int{}
    var v9 float32 = (((0.1 - 0.1) / (1e10 - float32(-2.25))) - (0.1 * 1.5))
    v10 := GenS1{F0: 0, F1: -9223372036854775808}
    var v11 uint32 = 818
    for fi12 := 0; fi12 < 5; fi12++ {
        v3 -= v3
        _ = fi12
    }
    var v13 string = ""
    if (1 <= 238) {
        for fi14 := 0; fi14 < 12; fi14++ {
            fmt.Println(fi14, v3, v3)
            _ = len(v1)
            _ = fi14
        }
    } else {
        if (!(false || false)) {
            v15 := v1[128]
            _ = v15
        }
        v11 += (^(v11))
        for fi16 := 0; fi16 < 8; fi16++ {
            v4 = append(v4, fi16)
            _ = fi16
        }
    }
    fmt.Println(v11, v3)
    v3 -= 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
