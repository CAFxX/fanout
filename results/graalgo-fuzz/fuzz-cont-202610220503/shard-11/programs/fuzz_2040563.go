package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
func genFn0(p0 uint, p1 bool) uint32 {
    v7 := GenS1{F0: 42, F1: 18446744073709551615}
    var v8 int16 = 597
    v8 -= v8
    var v9 int = 127
    v10 := []int{v9}
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 0
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 float64 = 3.14159
    v3 := make(map[int]bool)
    var v4 string = ""
    var v5 int16 = 32767
    v6 := make(map[string]int)
    v4 = "a"
    for fi11 := 0; fi11 < 9; fi11++ {
        for fi12 := 0; fi12 < 2; fi12++ {
            v5 += v5
            v13 := GenS1{F0: 2, F1: (v1 & v1)}
            _ = fi12
            _ = v13
        }
        _ = fi11
    }
    v1 -= (v1 ^ (v1 * 4294967295))
    v1 = (^(v1))
    v1 = 4294967295
    if (32767 > 1) {
        v4 = "foo bar"
    }
    v2 += v2
    v2 = (v2 - v2)
    v14 := GenS1{F0: 113, F1: (v1 - v1)}
    if true {
        v5 = (v5 * v5)
    }
    switch 91 {
    case -648:
        if (true || ((v4 + "") != "\n\t")) {
            v15 := GenS1{F0: 2, F1: v1}
            var v16 int32 = 0
            _ = v15
            _ = v16
        }
    case 9223372036854775807:
        var v17 uint = 286
        _ = v17
    case 127:
        var v18 float32 = ((1e10 - 3.14159) - float32(v2))
        _ = v18
    }
    v19 := v3[498]
    v1 = (v1 >> ((714) & 63))
    v19 = (v19 || ((v1 / ((v1) | 1)) == (v1 / ((v1) | 1))))
    v5 -= (v5 + v5)
    var v20 string = "hello"
    for fi21 := 0; fi21 < 4; fi21++ {
        for fi22 := 0; fi22 < 15; fi22++ {
            v2 = float64(0.0)
            _ = fi22
        }
        _ = fi21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
