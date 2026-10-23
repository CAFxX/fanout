package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 string }
func genFn0(p0 uint64, p1 int) string {
    v7 := GenS1{F0: (false && (42 >= 2147483647)), F1: (^(p0)), F2: ("" + ("\n\t" + "\n\t"))}
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    _ = v7
    return "hello"
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int = 9223372036854775807
    v3 := []string{"\n\t"}
    v4 := []uint32{}
    v5 := make(map[string]uint32)
    v6 := make(map[uint32]uint32)
    var v8 uint32 = 65535
    v8 = v8
    v2 = -539
    v2 += (v2 / ((127) | 1))
    switch v2 {
    case 127:
        for fi9 := 0; fi9 < 17; fi9++ {
            var v10 int16 = 42
            v10 = (v10 << ((952) & 63))
            _ = fi9
            _ = v10
        }
    case 248:
        var v11 uint = 4294967295
        _ = v11
    }
    fmt.Println(v8)
    for fi12 := 0; fi12 < 18; fi12++ {
        if len(v3) > 0 {
            v3[0] = "\n\t"
        }
        _ = fi12
    }
    v2 += -974
    v13 := GenS1{F0: (false || true), F1: 4294967295, F2: "a"}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
