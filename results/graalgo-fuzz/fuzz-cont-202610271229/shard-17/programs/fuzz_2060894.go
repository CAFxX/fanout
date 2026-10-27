package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 uint16; F1 uint32; F2 bool }
func genFn0(p0 uint8) int64 {
    v6 := []int{-258, 1}
    p0 += 130
    _ = p0
    _ = v6
    return 9223372036854775807
}

func main() {
    v1 := []string{("\n\t" + ""), ("" + ""), "foo bar", "foo bar"}
    var v2 uint8 = 42
    v3 := make(map[string]string)
    var v4 float32 = 1.5
    var v5 int16 = -1
    v4 -= v4
    fmt.Println(v2, v5, v2)
    for fi7 := 0; fi7 < 17; fi7++ {
        v4 = v4
        _ = fi7
    }
    v8 := GenS1{F0: (v5 ^ 1), F1: v2}
    v5 = 245
    v2 = 127
    v9 := GenS2{F0: 138, F1: 4294967295, F2: (-2.25 <= v4)}
    switch 9223372036854775807 {
    case -197:
        if len(v1) > 0 {
            v1[0] = ""
        }
    case 467:
        v4 = (v4 + (float32(0.0) + (v4 / v4)))
    }
    for fi10 := 0; fi10 < 3; fi10++ {
        if true {
            var v11 uint16 = 65535
            _ = v11
        }
        _ = fi10
    }
    if true {
        var v12 uint64 = 332
        _ = v12
    }
    var v13 float32 = v4
    v5 = (-(v5))
    var v14 uint8 = (v2 & 253)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
