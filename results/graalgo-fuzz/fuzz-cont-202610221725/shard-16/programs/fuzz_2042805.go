package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
func genFn0(p0 int, p1 string, p2 uint32) int32 {
    p1 = "a"
    p1 = ("" + p1)
    v4 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 42
}
func genFn1(p0 int64) bool {
    p0 -= p0
    var v5 string = ""
    if (18446744073709551615 <= 602) {
        v6 := GenS1{F0: 108, F1: -1}
        _ = v6
    }
    for fi7 := 0; fi7 < 15; fi7++ {
        p0 += 9223372036854775807
        switch 9223372036854775807 {
        case -150:
            var v8 int8 = 1
            _ = v8
        }
        _ = fi7
    }
    _ = p0
    _ = v5
    return (18446744073709551615 < 18446744073709551615)
}

func main() {
    v1 := []uint32{339, 4294967295, 2147483648}
    v2 := make(map[string]int)
    v3 := make(map[int]uint32)
    var v9 int32 = -2147483648
    var v10 int16 = -7
    for fi11 := 0; fi11 < 8; fi11++ {
        v1 = append(v1, 4294967295)
        _ = fi11
    }
    v10 = v10
    v1 = append(v1, 42)
    v12 := GenS1{F0: 65535, F1: (v10 - v10)}
    v10 -= 1
    var v13 int32 = (v9 - v9)
    v13 = v9
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v9
}
