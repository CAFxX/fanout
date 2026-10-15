package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 string }
func genFn0(p0 uint64, p1 int32, p2 string) int64 {
    v6 := GenS1{F0: -2147483648, F1: 0}
    for fi7 := 0; fi7 < 15; fi7++ {
        for fi8 := 0; fi8 < 12; fi8++ {
            var v9 float64 = (0.1 - 0.0)
            _ = fi8
            _ = v9
        }
        v10 := GenS1{F0: (p1 / ((p1) | 1)), F1: (fi7 / ((fi7) | 1))}
        _ = fi7
        _ = v10
    }
    switch (("hello" + ((p2 + "a") + ("" + p2))) + "hello") {
    case "\n\t":
        fmt.Println(p2, p0, p0)
    case "":
        v11 := GenS2{F0: p2}
        _ = v11
    case "a":
        p2 = "\n\t"
    }
    v12 := []int{32768}
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v6
    return 9223372036854775807
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[int]string)
    v3 := []int{-9223372036854775808, 879}
    var v4 uint16 = 32767
    var v5 int8 = 42
    switch 9223372036854775807 {
    case 42:
        var v13 int64 = -1
        _ = v13
    case 127:
        v14 := GenS2{F0: "foo bar"}
        _ = v14
    }
    fmt.Println(v4, v4)
    v3 = append(v3, -820)
    var v15 int32 = 810
    v16 := GenS2{F0: "foo bar"}
    v5 = (v5 >> ((646) & 63))
    v5 += -128
    fmt.Println(v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
