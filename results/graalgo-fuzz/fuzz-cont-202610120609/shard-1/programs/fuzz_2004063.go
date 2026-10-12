package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
func genFn0(p0 int64, p1 uint32) uint8 {
    var v6 int64 = p0
    var v7 int8 = 1
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 12
}

func main() {
    v1 := []string{"", "", "a"}
    v2 := make(map[string]bool)
    v3 := make(map[int]int)
    v4 := []string{}
    var v5 float32 = 0.0
    switch 574 {
    case 0:
        var v8 int32 = 0
        _ = v8
    case -527:
        for fi9 := 0; fi9 < 9; fi9++ {
            switch (fi9 % ((fi9) | 1)) {
            case 127:
                for fi10 := 0; fi10 < 18; fi10++ {
                    v5 += v5
                    v11 := GenS1{F0: 525, F1: 255}
                    _ = fi10
                    _ = v11
                }
            default:
                for fi12 := 0; fi12 < 5; fi12++ {
                    fmt.Println(fi12)
                    _ = fi12
                }
            }
            _ = fi9
        }
    case 1:
        switch "hello" {
        case "a":
            for fi13 := 0; fi13 < 10; fi13++ {
                v5 = (float32(-2.25) / v5)
                _ = fi13
            }
        case "hello":
            v14 := GenS1{F0: 1, F1: 1}
            _ = v14
        default:
            var v15 uint32 = 958
            _ = v15
        }
    default:
        var v16 float32 = (v5 / v5)
        _ = v16
    }
    delete(v2, (("foo bar" + "a") + ("\n\t" + "foo bar")))
    var v17 int8 = -128
    v18 := GenS1{F0: -680, F1: 42}
    v5 += 0.1
    _ = len(v2)
    fmt.Println(v17)
    var v19 uint = 0
    for fi20 := 0; fi20 < 11; fi20++ {
        for fi21 := 0; fi21 < 5; fi21++ {
            var v22 uint16 = 32767
            fmt.Println(v22, fi20, v19)
            fmt.Println(v17)
            fmt.Println(fi20, v17)
            _ = fi21
            _ = v22
        }
        v19 = (v19 * v19)
        v5 -= (v5 + v5)
        _ = fi20
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
