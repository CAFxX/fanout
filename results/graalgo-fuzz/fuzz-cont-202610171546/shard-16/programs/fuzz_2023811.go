package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }
func genFn0() int {
    var v5 int64 = 42
    for fi6 := 0; fi6 < 17; fi6++ {
        var v7 float32 = ((1.5 / (-2.25 + (1e10 - (1.5 * (float32(((0.0 / (1.5 - 3.14159)) + 3.14159)) / (1.5 + (1e10 / ((1e10 - 0.1) * 1.5)))))))) - 0.1)
        fmt.Println(v7, v7)
        _ = fi6
        _ = v7
    }
    _ = v5
    return 4294967295
}

func main() {
    var v1 uint64 = 406
    var v2 int32 = -2147483648
    v3 := []string{"hello", "a"}
    var v4 int64 = -699
    v3 = append(v3, "a")
    for fi8 := 0; fi8 < 7; fi8++ {
        v2 = (v2 % ((v2) | 1))
        _ = fi8
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v4, v1)
    v2 -= (^(v2))
    v4 = 9223372036854775807
    if false {
        var v9 int32 = v2
        _ = v9
    }
    v10 := GenS1{F0: (!(0.0 != 3.14159)), F1: 42}
    if (42 <= v2) {
        v2 += (v2 << ((65535) & 63))
        v11 := GenS1{F0: (true || true), F1: int(v2)}
        _ = v11
    }
    _ = len(v3)
    var v12 int8 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
    _ = v4
}
