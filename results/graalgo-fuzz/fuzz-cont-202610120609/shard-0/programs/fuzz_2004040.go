package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 uint64 }
type GenS2 struct { F0 bool; F1 uint32; F2 uint16 }
func genFn0(p0 int16) int32 {
    v4 := GenS1{F0: 42, F1: -1, F2: 18}
    var v5 int64 = 42
    v5 += v5
    _ = p0
    _ = v4
    _ = v5
    return -245
}

func main() {
    var v1 uint64 = 1
    v2 := make(map[string]int)
    v3 := make(map[int]string)
    v3[0] = ("" + ("hello" + "\n\t"))
    var v6 int = -9223372036854775808
    v7 := []int{v6}
    for fi8 := 0; fi8 < 2; fi8++ {
        var v9 uint = 0
        _ = fi8
        _ = v9
    }
    fmt.Println(v6, v6)
    if (float32(3.14159) < 0.0) {
        var v10 int = (v6 - 128)
        if false {
            for fi11 := 0; fi11 < 2; fi11++ {
                if false {
                    v12 := GenS2{F0: ("hello" == "\n\t"), F1: 2147483647, F2: 2}
                    var v13 float64 = 3.14159
                    v3[v6] = ""
                    _ = v12
                    _ = v13
                }
                v10 = v10
                _ = fi11
            }
            v14 := GenS2{F0: false, F1: 1, F2: 42}
            _ = v14
        }
        v6 += v10
        _ = v10
    }
    v7 = append(v7, v6)
    for fi15 := 0; fi15 < 7; fi15++ {
        switch "a" {
        case "hello":
            var v16 uint = 590
            _ = v16
        case "a":
            v1 -= (v1 % ((v1) | 1))
        }
        _ = fi15
    }
    v6 = v6
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v6
    _ = v7
}
