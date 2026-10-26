package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 uint }
func genFn0(p0 uint32) uint64 {
    var v7 uint32 = p0
    var v8 int64 = -1
    _ = p0
    _ = v7
    _ = v8
    return 4294967295
}
func genFn1(p0 uint8) int32 {
    switch 127 {
    case 4294967295:
        if (!false) {
            fmt.Println(p0)
        }
    }
    fmt.Println(p0)
    p0 += 42
    p0 -= p0
    _ = p0
    return -896
}

func main() {
    v1 := []string{"", ("\n\t" + ""), "\n\t", ""}
    v2 := make(map[int]uint32)
    var v3 uint = 18446744073709551615
    v4 := []int{-639, 32768, 127, 4294967295}
    var v5 uint32 = 0
    var v6 bool = (888 <= uint64(v3))
    v5 += v5
    for fi9 := 0; fi9 < 12; fi9++ {
        v6 = false
        _ = fi9
    }
    for fi10 := 0; fi10 < 14; fi10++ {
        var v11 int64 = -947
        if v6 {
            var v12 float64 = 1.5
            if len(v4) > 0 {
                v4[0] = fi10
            }
            _ = v12
        }
        v13 := GenS1{F0: (!v6), F1: ("a" + "\n\t"), F2: v3}
        _ = fi10
        _ = v11
        _ = v13
    }
    v14 := v2[9223372036854775807]
    for fi15 := 0; fi15 < 19; fi15++ {
        var v16 int8 = -128
        v17 := GenS1{F0: (!(-9223372036854775808 < -9223372036854775808)), F1: "\n\t", F2: (v3 | (v3 + 42))}
        _ = fi15
        _ = v16
        _ = v17
    }
    v18 := GenS1{F0: (52 == 2), F1: "\n\t", F2: v3}
    v19 := GenS1{F0: v6, F1: "hello", F2: v3}
    v3 += v3
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
