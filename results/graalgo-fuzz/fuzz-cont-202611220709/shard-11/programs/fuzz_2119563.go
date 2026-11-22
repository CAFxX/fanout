package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
func genFn0() uint {
    var v4 int64 = 42
    v5 := GenS1{F0: -273, F1: -128}
    v4 -= (-(v4))
    _ = v4
    _ = v5
    return 106
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[string]bool)
    var v3 bool = false
    v6 := []int{709, 9223372036854775807, 2147483647, 4294967295}
    v7 := v1["a"]
    v2["foo bar"] = (true || v3)
    v8 := GenS1{F0: 1, F1: 127}
    v7 -= -34
    if (157 <= 92) {
        if (0 >= -9223372036854775808) {
            var v9 int64 = 42
            v7 += v7
            for fi10 := 0; fi10 < 13; fi10++ {
                v7 += fi10
                _ = fi10
            }
            _ = v9
        } else {
            var v11 int64 = 710
            _ = v11
        }
        if ((0.0 != -2.25) && (!v3)) {
            v7 = (v7 | -2147483648)
        } else {
            fmt.Println(v7, v3)
            _ = len(v1)
            v3 = false
        }
        for fi12 := 0; fi12 < 9; fi12++ {
            v13 := GenS1{F0: 909, F1: -128}
            _ = fi12
            _ = v13
        }
    } else {
        fmt.Println(v3, v3)
    }
    switch v7 {
    case 297:
        v14 := v1["hello"]
        _ = v14
    case 0:
        if (3.14159 >= 0.0) {
            var v15 float32 = 3.14159
            if v3 {
                var v16 uint16 = 42
                _ = v16
            }
            v3 = (false || true)
            _ = v15
        } else {
            var v17 uint32 = 623
            _ = v17
        }
    }
    if false {
        v3 = v3
    }
    fmt.Println(v7, v7)
    switch -9223372036854775808 {
    case -1:
        v18 := GenS1{F0: -2, F1: -31}
        _ = v18
    case 4294967295:
        delete(v1, "\n\t")
    case 0:
        v7 -= -658
    }
    var v19 uint64 = 116
    if false {
        v19 = (v19 ^ (v19 * (v19 / ((v19) | 1))))
        var v20 int32 = -60
        fmt.Println(v3, v19)
        _ = v20
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
}
