package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
func genFn0(p0 string, p1 int, p2 uint16) uint32 {
    var v4 bool = true
    p1 = p1
    for fi5 := 0; fi5 < 20; fi5++ {
        v6 := GenS1{F0: (-(p1)), F1: 18446744073709551615}
        var v7 int32 = 24
        _ = fi5
        _ = v6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 0
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[string]int)
    var v3 bool = false
    var v8 int8 = -122
    fmt.Println(v3, v8)
    var v9 int32 = 843
    v3 = (!(((1e10 * (-2.25 + ((float64(0.1) * 3.14159) * float64(1.5)))) - float64(0.0)) != 1e10))
    if ((v3 && (!(821 > 134))) && v3) {
        var v10 uint16 = 42
        var v11 uint64 = 42
        if false {
            for fi12 := 0; fi12 < 13; fi12++ {
                delete(v1, fi12)
                if v3 {
                    for fi13 := 0; fi13 < 15; fi13++ {
                        if v3 {
                            v3 = ((fi13 ^ fi12) != (fi12 - (fi12 % (((fi12 | (^(fi13)))) | 1))))
                            fmt.Println(v8, v3, v9)
                            var v14 bool = (0.0 < 3.14159)
                            _ = v14
                        }
                        _ = fi13
                    }
                    var v15 uint64 = 42
                    for fi16 := 0; fi16 < 1; fi16++ {
                        var v17 int16 = 559
                        var v18 float32 = float32(1e10)
                        v19 := v2[("hello" + "a")]
                        v15 += v11
                        _ = fi16
                        _ = v17
                        _ = v18
                        _ = v19
                    }
                    _ = v15
                } else {
                    v3 = ((fi12 | 1) != fi12)
                    v20 := GenS1{F0: fi12, F1: 2}
                    _ = v20
                }
                _ = fi12
            }
        }
        _ = v10
        _ = v11
    }
    v9 = 2147483647
    fmt.Println(v9, v9, v3)
    _ = len(v2)
    v21 := []int{-1, 2147483647, 4294967295}
    v8 += v8
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v8
    _ = v9
}
