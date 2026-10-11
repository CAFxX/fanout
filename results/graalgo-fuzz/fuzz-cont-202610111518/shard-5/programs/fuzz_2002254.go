package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 uint16 }
type GenS2 struct { F0 int8; F1 int64; F2 string }
func genFn0(p0 bool) string {
    v4 := []int{-554, 32768, 451}
    _ = len(v4)
    _ = p0
    _ = v4
    return "hello"
}

func main() {
    var v1 float64 = (float64(0.1) - (1e10 - 3.14159))
    v2 := make(map[uint32]string)
    v3 := make(map[string]int)
    v5 := GenS1{F0: 2, F1: -1, F2: 891}
    fmt.Println(v1)
    v6 := []int{-1, 128}
    delete(v3, "a")
    switch -56 {
    case 32767:
        var v7 uint = 4294967295
        _ = v7
    case 2147483647:
        v1 += (v1 * v1)
    case -1:
        for fi8 := 0; fi8 < 13; fi8++ {
            v1 += v1
            if len(v6) > 0 {
                v6[0] = 872
            }
            v9 := v2[0]
            _ = fi8
            _ = v9
        }
    default:
        var v10 uint16 = 32767
        _ = v10
    }
    var v11 int16 = -1
    v12 := GenS2{F0: 56, F1: -9223372036854775808, F2: "foo bar"}
    var v13 int64 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v5
    _ = v6
}
