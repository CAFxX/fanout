package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
func genFn0(p0 int64) int {
    fmt.Println(p0)
    var v6 uint16 = 1
    _ = p0
    _ = v6
    return -552
}

func main() {
    var v1 uint16 = 45
    v2 := make(map[string]bool)
    var v3 float32 = 1e10
    v4 := make(map[string]int)
    v5 := make(map[uint32]int)
    switch "a" {
    case "\n\t":
        v1 -= (v1 >> ((1) & 63))
    default:
        v3 -= ((0.1 * float32((-2.25 + -2.25))) - (v3 / v3))
    }
    delete(v4, "\n\t")
    v7 := GenS1{F0: 0, F1: -553}
    for fi8 := 0; fi8 < 9; fi8++ {
        var v9 uint64 = 151
        _ = fi8
        _ = v9
    }
    v10 := []int{4294967295, 128}
    for fi11 := 0; fi11 < 1; fi11++ {
        for fi12 := 0; fi12 < 10; fi12++ {
            v13 := GenS1{F0: 735, F1: 42}
            var v14 uint = 2
            v1 += 65535
            _ = fi12
            _ = v13
            _ = v14
        }
        _ = fi11
    }
    switch "" {
    case "":
        switch 4294967295 {
        case -480:
            fmt.Println(v1, v3)
        case 42:
            var v15 string = ("\n\t" + "a")
            _ = v15
        case 9223372036854775807:
            for fi16 := 0; fi16 < 8; fi16++ {
                var v17 int8 = 88
                for fi18 := 0; fi18 < 15; fi18++ {
                    fmt.Println(fi18, v3, v17)
                    _ = fi18
                }
                _ = fi16
                _ = v17
            }
        }
    }
    v3 += 0.1
    v1 = (v1 / ((65535) | 1))
    _ = len(v5)
    v19 := v2["\n\t"]
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
