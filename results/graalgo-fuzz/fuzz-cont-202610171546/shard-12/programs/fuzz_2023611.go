package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 int }
func genFn0(p0 uint64, p1 uint32) bool {
    p0 = (p0 - (p0 ^ 42))
    var v6 float32 = (1e10 + ((1e10 + (-2.25 / 1.0)) * 1e10))
    _ = p0
    _ = p1
    _ = v6
    return true
}
func genFn1(p0 int8, p1 uint8) uint8 {
    p1 += p1
    var v7 float32 = 1.5
    var v8 int16 = 271
    v9 := GenS1{F0: v8, F1: "foo bar", F2: 668}
    for fi10 := 0; fi10 < 7; fi10++ {
        v11 := []int{(-(fi10)), (fi10 << ((255) & 63)), fi10}
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return (^(p1))
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[uint32]uint32)
    v3 := make(map[int]string)
    var v4 int = -2147483648
    var v5 float64 = 0.1
    v12 := []int{32767, 2147483647}
    v4 -= v4
    v4 = v4
    if false {
        if true {
            for fi13 := 0; fi13 < 8; fi13++ {
                v5 = v5
                _ = fi13
            }
        }
        if ((-1 < -124) && (((("a" + "\n\t") + "a") + "hello") != "foo bar")) {
            if len(v12) > 0 {
                v12[0] = (^(v4))
            }
            v5 = 1.5
        } else {
            _ = len(v12)
            fmt.Println(v5)
            v14 := v1["hello"]
            _ = v14
        }
        var v15 int8 = 100
        _ = v15
    }
    v5 -= 1e10
    switch v4 {
    case 128:
        var v16 uint16 = 660
        _ = v16
    default:
        _ = len(v12)
    }
    v4 -= 559
    v4 += v4
    v4 = v4
    for fi17 := 0; fi17 < 5; fi17++ {
        v5 += 1.5
        for fi18 := 0; fi18 < 12; fi18++ {
            if (!(196 == 42)) {
                var v19 uint64 = 0
                v20 := GenS1{F0: 0, F1: "hello", F2: v4}
                _ = v19
                _ = v20
            }
            _ = fi18
        }
        _ = fi17
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
