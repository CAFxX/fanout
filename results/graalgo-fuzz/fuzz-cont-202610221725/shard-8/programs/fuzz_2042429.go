package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0(p0 int16, p1 int64) int32 {
    v5 := []int{9223372036854775807, 4294967295}
    p0 += (p0 + (p0 + (p0 / (((p0 % ((p0) | 1))) | 1))))
    var v6 int64 = (p1 + p1)
    var v7 int16 = (p0 | p0)
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return -442
}

func main() {
    v1 := []int{4294967295, -949, 2147483647}
    var v2 uint8 = 127
    var v3 uint16 = 555
    var v4 int64 = -1
    v2 = v2
    var v8 float64 = 3.14159
    if len(v1) > 0 {
        v1[0] = 128
    }
    var v9 uint32 = 42
    fmt.Println(v8, v4)
    if false {
        if ("" != ("hello" + ("foo bar" + ("" + "\n\t")))) {
            v4 -= v4
        }
        var v10 int32 = -2147483648
        _ = v10
    }
    if (0 >= -1) {
        for fi11 := 0; fi11 < 17; fi11++ {
            v2 = 158
            _ = fi11
        }
        for fi12 := 0; fi12 < 18; fi12++ {
            if len(v1) > 0 {
                v1[0] = 769
            }
            _ = fi12
        }
    } else {
        v3 = (v3 & (v3 & 65535))
    }
    v1 = append(v1, 0)
    for fi13 := 0; fi13 < 20; fi13++ {
        v14 := GenS1{F0: (v3 | v3)}
        var v15 bool = (1 <= 1)
        _ = fi13
        _ = v14
        _ = v15
    }
    v3 += 518
    var v16 int = 4294967295
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
