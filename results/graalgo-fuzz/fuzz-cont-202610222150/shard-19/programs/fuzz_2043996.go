package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
func genFn0(p0 int, p1 string) int8 {
    switch "a" {
    case "":
        p0 -= 127
    }
    var v4 int = (p0 ^ (p0 | (^(p0))))
    var v5 int = p0
    v6 := []int{}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return -128
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint64 = 846
    v3 := make(map[string]int)
    switch "a" {
    case "a":
        for fi7 := 0; fi7 < 13; fi7++ {
            v8 := GenS1{F0: 129, F1: -9223372036854775808}
            v9 := GenS1{F0: 0, F1: 42}
            _ = fi7
            _ = v8
            _ = v9
        }
    }
    v10 := GenS1{F0: 127, F1: -9223372036854775808}
    var v11 float32 = -2.25
    v12 := GenS1{F0: 255, F1: -427}
    var v13 float32 = v11
    var v14 float32 = (v11 - ((v13 + (-2.25 * 1.5)) + v11))
    v15 := GenS1{F0: 128, F1: 9223372036854775807}
    if true {
        for fi16 := 0; fi16 < 16; fi16++ {
            _ = len(v1)
            _ = fi16
        }
    } else {
        fmt.Println(v11, v2)
    }
    var v17 uint = 18446744073709551615
    v18 := GenS1{F0: 212, F1: 1}
    var v19 int64 = 254
    fmt.Println(v19, v19, v2)
    v17 -= v17
    v19 -= (v19 / (((v19 >> (((v17 & 2)) & 63))) | 1))
    v20 := []int{4294967295, -808, -955, 781}
    if false {
        if len(v20) >= 2 {
            v20 = v20[1:len(v20)]
        }
        v20 = append(v20, 426)
    } else {
        v11 = 1.5
        v21 := GenS1{F0: 0, F1: (v19 / ((v19) | 1))}
        _ = v21
    }
    for fi22 := 0; fi22 < 11; fi22++ {
        for fi23 := 0; fi23 < 1; fi23++ {
            v13 = -2.25
            if false {
                v24 := GenS1{F0: 78, F1: v19}
                _ = v24
            }
            _ = fi23
        }
        fmt.Println(fi22)
        _ = fi22
    }
    v25 := GenS1{F0: 42, F1: (v19 % (((v19 >> ((873) & 63))) | 1))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v25
    _ = v3
}
