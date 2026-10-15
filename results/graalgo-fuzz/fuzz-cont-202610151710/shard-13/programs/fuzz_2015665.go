package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool; F1 string }
func genFn0() int8 {
    for fi6 := 0; fi6 < 20; fi6++ {
        var v7 bool = (((0 >= 0) && (!false)) && (((float32(1e10) - (1.5 - (0.0 - 0.0))) == (-2.25 - (float32(-2.25) - float32(0.1)))) || ((110 == 1) || (!true))))
        for fi8 := 0; fi8 < 2; fi8++ {
            v9 := []int{}
            var v10 int64 = 42
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = fi6
        _ = v7
    }
    v11 := []int{2147483647}
    _ = v11
    return -21
}

func main() {
    var v1 uint = 163
    v2 := []string{"\n\t"}
    v3 := make(map[int]int)
    v4 := []string{(("\n\t" + (("\n\t" + "") + "a")) + "\n\t"), ("a" + ("" + "hello"))}
    var v5 uint = (v1 % ((v1) | 1))
    fmt.Println(v5, v1)
    v12 := GenS1{F0: 1000}
    v5 = (v5 - 0)
    if (0 == 32767) {
        for fi13 := 0; fi13 < 1; fi13++ {
            _ = len(v4)
            _ = fi13
        }
        for fi14 := 0; fi14 < 7; fi14++ {
            var v15 float32 = (1.5 - (float32(((3.14159 / (0.0 + (0.1 * 1.5))) + 0.1)) + float32((-2.25 * float64((0.1 - 0.0))))))
            fmt.Println(v1, v1, v5)
            if (477 > 543) {
                _ = len(v3)
                _ = len(v4)
                v16 := GenS1{F0: -1}
                _ = v16
            } else {
                if len(v2) > 0 {
                    v2[0] = "a"
                }
            }
            _ = fi14
            _ = v15
        }
        if true {
            var v17 uint32 = 0
            _ = v17
        } else {
            v1 -= 1
            fmt.Println(v5, v1)
            var v18 uint64 = 18446744073709551615
            _ = v18
        }
    } else {
        v19 := GenS2{F0: (0 <= 42), F1: "a"}
        _ = v19
    }
    v20 := GenS1{F0: 1000}
    v1 += v1
    if (!false) {
        var v21 uint32 = 42
        var v22 int = -337
        _ = v21
        _ = v22
    }
    var v23 uint8 = 0
    var v24 uint16 = 1
    v24 = 191
    var v25 uint64 = 1
    var v26 uint16 = v24
    var v27 uint8 = v23
    v28 := v3[42]
    v29 := GenS1{F0: 32767}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
