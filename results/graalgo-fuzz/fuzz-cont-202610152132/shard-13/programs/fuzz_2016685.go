package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }

func main() {
    var v1 string = "foo bar"
    var v2 int = 1
    v3 := make(map[int]string)
    var v4 float32 = 0.1
    v5 := make(map[int]bool)
    v6 := make(map[int]int)
    for fi7 := 0; fi7 < 16; fi7++ {
        v4 -= (((v4 / float32(0.1)) / float32(1.5)) / float32(0.1))
        _ = fi7
    }
    for fi8 := 0; fi8 < 8; fi8++ {
        v2 = -952
        v2 += v2
        _ = fi8
    }
    for fi9 := 0; fi9 < 20; fi9++ {
        v1 = "foo bar"
        for fi10 := 0; fi10 < 18; fi10++ {
            var v11 int64 = 0
            v12 := GenS1{F0: 42, F1: 83}
            if (true && (2147483647 != 483)) {
                v13 := []int{(^(v2)), (-(v2)), (fi10 * (fi10 + fi10)), 127}
                var v14 bool = false
                if len(v13) > 0 {
                    v13[0] = fi9
                }
                _ = v13
                _ = v14
            }
            _ = fi10
            _ = v11
            _ = v12
        }
        if (true || false) {
            v15 := GenS1{F0: 1, F1: -128}
            var v16 string = (v1 + "a")
            v2 += fi9
            _ = v15
            _ = v16
        }
        _ = fi9
    }
    v17 := []int{(v2 >> ((2) & 63)), v2, (v2 & v2)}
    if false {
        v3[(v2 | v2)] = v1
        _ = len(v17)
    }
    var v18 int32 = -2147483648
    v1 = (v1 + v1)
    v2 = v2
    v1 = "hello"
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
