package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
type GenS2 struct { F0 int32; F1 int64 }
type GenS3 struct { F0 uint }
func genFn0(p0 int64, p1 int, p2 uint16) bool {
    v7 := []int{p1, (^(p1))}
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    if true {
        v7 = append(v7, p1)
    } else {
        v8 := GenS1{F0: 1, F1: -89}
        var v9 uint = 0
        p1 = p1
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return (-1 != -30)
}

func main() {
    var v1 float32 = ((1e10 / 1.0) - (float32(1e10) - 1e10))
    v2 := make(map[uint32]string)
    v3 := []string{"foo bar"}
    v4 := make(map[string]string)
    var v5 int32 = 396
    var v6 int32 = 2147483647
    if (0 != 40) {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    } else {
        var v10 float64 = 1e10
        if ((v6 << ((2) & 63)) < v5) {
            var v11 int32 = (v6 % ((v6) | 1))
            _ = v11
        } else {
            v12 := v2[0]
            v13 := GenS2{F0: 695, F1: -1}
            _ = v12
            _ = v13
        }
        var v14 uint32 = 65535
        _ = v10
        _ = v14
    }
    v15 := GenS3{F0: 65535}
    for fi16 := 0; fi16 < 6; fi16++ {
        v2[2147483647] = "hello"
        var v17 uint16 = 1
        _ = fi16
        _ = v17
    }
    var v18 float64 = float64(v1)
    for fi19 := 0; fi19 < 6; fi19++ {
        v20 := GenS1{F0: 65535, F1: 0}
        v6 = v6
        _ = fi19
        _ = v20
    }
    v5 = v5
    var v21 int8 = -128
    v21 = -75
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
