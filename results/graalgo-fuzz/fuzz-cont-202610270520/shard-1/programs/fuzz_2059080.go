package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
func genFn0(p0 uint, p1 int16) uint {
    for fi4 := 0; fi4 < 1; fi4++ {
        for fi5 := 0; fi5 < 1; fi5++ {
            var v6 uint = p0
            _ = fi5
            _ = v6
        }
        v7 := []int{(^(fi4)), (^(fi4))}
        _ = fi4
        _ = v7
    }
    var v8 int16 = (p1 | p1)
    _ = p0
    _ = p1
    _ = v8
    return (p0 % (((^(p0))) | 1))
}
func genFn1(p0 int8, p1 int8) int16 {
    if true {
        v9 := GenS1{F0: -32768, F1: false}
        p1 = 125
        _ = v9
    }
    p0 = (p1 << ((255) & 63))
    v10 := []int{4294967295}
    for fi11 := 0; fi11 < 7; fi11++ {
        var v12 uint = 65535
        if ("a" > "a") {
            if (!false) {
                fmt.Println(fi11, p0, v12)
                fmt.Println(p0)
            } else {
                var v13 uint32 = 65535
                if len(v10) > 0 {
                    v10[0] = fi11
                }
                _ = v13
            }
        }
        _ = fi11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v10
    return 203
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]int)
    var v3 bool = true
    v3 = v3
    v3 = ((65535 == 518) || (1000 < 0))
    fmt.Println(v3)
    var v14 float64 = (float64(1.5) + -2.25)
    fmt.Println(v14, v14)
    v3 = false
    if false {
        var v15 int = 32768
        v3 = v3
        _ = v15
    } else {
        v16 := GenS1{F0: 0, F1: v3}
        fmt.Println(v14, v3)
        _ = v16
    }
    v14 = (((v14 + v14) - 0.0) + 1.5)
    v17 := []int{4294967295, 953, -262, -443}
    var v18 float32 = (0.1 / (1.5 / (float32(float64(1.5)) / 1e10)))
    var v19 int32 = 915
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}
