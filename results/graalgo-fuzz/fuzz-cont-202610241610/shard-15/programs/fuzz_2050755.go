package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 string }
func genFn0(p0 int32) int64 {
    p0 = p0
    p0 = (p0 - p0)
    for fi7 := 0; fi7 < 8; fi7++ {
        v8 := []int{(fi7 >> ((381) & 63)), fi7, 9223372036854775807, 9223372036854775807}
        p0 = 696
        _ = fi7
        _ = v8
    }
    _ = p0
    return 230
}

func main() {
    var v1 int32 = -1
    v2 := make(map[int]string)
    var v3 string = "a"
    var v4 bool = (42 == 0)
    v5 := make(map[uint32]bool)
    v6 := make(map[string]bool)
    var v9 uint64 = 2
    v10 := []int{4294967295, 127}
    var v11 uint = 2
    v1 -= (v1 & 311)
    var v12 int64 = 579
    var v13 int32 = v1
    if (820 >= 42) {
        var v14 int16 = 196
        var v15 uint32 = 210
        _ = v14
        _ = v15
    } else {
        switch v3 {
        case "hello":
            v10 = append(v10, 1)
        }
        var v16 uint8 = 128
        if (65535 == 553) {
            if ((v9 % ((4294967295) | 1)) <= (^(v9))) {
                var v17 bool = (v4 && true)
                _ = v17
            } else {
                v18 := GenS1{F0: -169, F1: 2147483647, F2: ("hello" + "hello")}
                for fi19 := 0; fi19 < 4; fi19++ {
                    var v20 float64 = (3.14159 / float64(1.5))
                    var v21 uint64 = v9
                    _ = fi19
                    _ = v20
                    _ = v21
                }
                _ = v18
            }
        }
        _ = v16
    }
    var v22 int16 = 42
    fmt.Println(v3, v1, v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
