package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 bool }
func genFn0(p0 int16, p1 bool, p2 int16) uint32 {
    v4 := GenS1{F0: 42, F1: 652, F2: (p1 && (p0 <= (p2 | 601)))}
    var v5 int8 = 42
    p0 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 2147483647
}

func main() {
    var v1 int8 = -97
    v2 := make(map[int]string)
    v3 := []int{42, 116, 32767, 127}
    switch 1 {
    case -1:
        v1 += v1
    case -2147483648:
        delete(v2, 2147483647)
    }
    switch -684 {
    case 2147483647:
        if (true || true) {
            for fi6 := 0; fi6 < 16; fi6++ {
                v1 = (v1 * v1)
                _ = fi6
            }
            for fi7 := 0; fi7 < 5; fi7++ {
                var v8 bool = true
                _ = fi7
                _ = v8
            }
            for fi9 := 0; fi9 < 6; fi9++ {
                var v10 int = (-(fi9))
                fmt.Println(fi9, v1)
                v10 = 2147483647
                _ = fi9
                _ = v10
            }
        }
    }
    v1 -= -128
    fmt.Println(v1)
    v1 = 0
    fmt.Println(v1)
    fmt.Println(v1)
    v1 = 9
    v1 = (v1 & (v1 % ((v1) | 1)))
    var v11 uint = 42
    for fi12 := 0; fi12 < 4; fi12++ {
        switch "" {
        case "foo bar":
            v13 := GenS1{F0: 4294967295, F1: v11, F2: false}
            _ = v13
        case "\n\t":
            var v14 uint32 = 2
            _ = v14
        case "hello":
            if len(v3) > 0 {
                v3[0] = (fi12 ^ 103)
            }
        }
        if (922 < 4294967295) {
            v1 -= (v1 | 1)
            v15 := GenS1{F0: 18446744073709551615, F1: v11, F2: (((0.1 + 1.5) * ((((-2.25 - (-2.25 * 0.1)) / (3.14159 * (0.1 - ((1.5 - ((1e10 * 0.1) + 0.0)) + 1e10)))) - 3.14159) * (3.14159 * 1.5))) != -2.25)}
            var v16 uint32 = 926
            _ = v15
            _ = v16
        }
        _ = fi12
    }
    if true {
        var v17 uint32 = 947
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
}
