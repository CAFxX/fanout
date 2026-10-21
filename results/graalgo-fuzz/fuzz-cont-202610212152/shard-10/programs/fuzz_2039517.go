package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 int64 }
func genFn0(p0 int32, p1 int8, p2 bool) string {
    switch "hello" {
    case "\n\t":
        v5 := []int{4294967295, -9223372036854775808, 758}
        _ = v5
    }
    v6 := GenS1{F0: p0, F1: -32768, F2: 9223372036854775807}
    var v7 int8 = p1
    p1 -= (-(v7))
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return "\n\t"
}

func main() {
    v1 := make(map[int]bool)
    var v2 int32 = 0
    v3 := make(map[uint32]string)
    v4 := []string{("" + "a"), "foo bar", "\n\t", "hello"}
    var v8 int = 42
    var v9 uint16 = 65535
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v10 uint8 = 42
    for fi11 := 0; fi11 < 9; fi11++ {
        v2 += v2
        v10 = v10
        var v12 bool = false
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 14; fi13++ {
        var v14 int8 = -105
        _ = fi13
        _ = v14
    }
    var v15 int = (v8 - (v8 * 2147483647))
    var v16 int32 = (v2 / ((v2) | 1))
    v1[(v8 - (^(v8)))] = false
    fmt.Println(v16)
    var v17 string = "hello"
    v18 := GenS1{F0: (v16 * 0), F1: int16(v16), F2: -1}
    v8 -= v15
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
