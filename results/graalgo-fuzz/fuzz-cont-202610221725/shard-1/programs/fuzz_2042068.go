package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 int64 }

func main() {
    var v1 int16 = 1000
    var v2 int = 1
    v3 := []uint32{2147483648, 42, 4294967295}
    v4 := make(map[string]string)
    var v5 bool = (2147483647 <= 117)
    var v6 uint32 = 2147483648
    v6 -= 65535
    if v5 {
        v2 += 128
        _ = len(v4)
        if (65535 > 532) {
            fmt.Println(v1, v2, v6)
        }
    } else {
        var v7 int32 = -838
        v8 := GenS1{F0: 1, F1: "foo bar", F2: -1}
        _ = v7
        _ = v8
    }
    if len(v3) > 0 {
        v3[0] = (v6 << ((1) & 63))
    }
    v1 += (-(v1))
    v1 += (v1 << ((1) & 63))
    v6 -= (v6 - (v6 - (v6 | (^(v6)))))
    var v9 float32 = ((1.5 * 3.14159) / 1.5)
    v6 -= (v6 << ((215) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
