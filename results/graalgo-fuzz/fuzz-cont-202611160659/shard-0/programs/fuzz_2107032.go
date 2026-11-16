package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string; F1 int16 }
func genFn0() int32 {
    var v7 bool = false
    if (-2.25 > (0.0 + 1e10)) {
        switch 9223372036854775807 {
        case 32768:
            for fi8 := 0; fi8 < 14; fi8++ {
                v7 = v7
                _ = fi8
            }
        }
        v9 := GenS1{F0: 32768}
        _ = v9
    } else {
        v10 := []int{}
        var v11 uint8 = 250
        if (!v7) {
            if (-79 >= -440) {
                v7 = (false && (0.0 > (1.5 * 0.0)))
            } else {
                var v12 int16 = 193
                var v13 uint64 = 0
                fmt.Println(v12, v7, v12)
                _ = v12
                _ = v13
            }
            var v14 int = -456
            fmt.Println(v14, v7, v7)
            _ = v14
        }
        _ = v10
        _ = v11
    }
    v15 := GenS1{F0: 4294967295}
    _ = v15
    _ = v7
    return 2147483647
}
func genFn1(p0 uint, p1 string) int32 {
    var v16 float64 = 3.14159
    var v17 uint32 = 462
    _ = p0
    _ = p1
    _ = v16
    _ = v17
    return 766
}

func main() {
    var v1 float32 = (1.5 + 0.1)
    v2 := make(map[int]int)
    var v3 float64 = 0.1
    v4 := make(map[string]int)
    var v5 int16 = -746
    v6 := []int{9223372036854775807, 4294967295, 1}
    v6 = append(v6, 784)
    v3 = 3.14159
    for fi18 := 0; fi18 < 13; fi18++ {
        v19 := GenS1{F0: fi18}
        _ = fi18
        _ = v19
    }
    if (-2.25 <= (3.14159 / v3)) {
        _ = len(v4)
    }
    v1 += 3.14159
    if true {
        v20 := GenS2{F0: "a", F1: -1}
        var v21 string = "a"
        _ = v20
        _ = v21
    }
    fmt.Println(v3, v1)
    var v22 string = "foo bar"
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
