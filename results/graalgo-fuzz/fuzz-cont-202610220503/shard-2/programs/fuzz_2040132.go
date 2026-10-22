package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 uint32 }
func genFn0() int64 {
    var v6 uint16 = 42
    if false {
        v7 := []int{}
        var v8 int = 128
        _ = v7
        _ = v8
    }
    v6 = 255
    _ = v6
    return 0
}
func genFn1(p0 uint8, p1 uint) int {
    p0 += p0
    if (42 > 0) {
        for fi9 := 0; fi9 < 10; fi9++ {
            p0 = p0
            fmt.Println(p0, p0, p1)
            _ = fi9
        }
    } else {
        v10 := GenS1{F0: true, F1: 0, F2: 891}
        p0 -= (p0 & p0)
        _ = v10
    }
    var v11 int16 = -32768
    var v12 uint = (^(p1))
    v13 := GenS1{F0: (1 >= 4294967295), F1: 2, F2: 0}
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    _ = v13
    return 9223372036854775807
}

func main() {
    var v1 int = 9223372036854775807
    v2 := make(map[int]string)
    var v3 int64 = 1
    v4 := make(map[string]string)
    var v5 int8 = 14
    var v14 float32 = float32(1.5)
    fmt.Println(v3)
    var v15 uint16 = 42
    v16 := GenS1{F0: true, F1: 4294967295, F2: 2147483648}
    v17 := []int{(v1 | genFn1(192, 4294967295)), (^(v1)), (v1 - 468)}
    v3 = genFn0()
    _ = len(v17)
    v14 += 1.5
    var v18 float64 = (0.1 * (((-2.25 + 0.0) / -2.25) / float64((((v14 / float32(3.14159)) + v14) + v14))))
    var v19 uint16 = 122
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
