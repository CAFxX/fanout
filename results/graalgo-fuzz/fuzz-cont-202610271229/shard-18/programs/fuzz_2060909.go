package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
func genFn0(p0 uint64) int64 {
    v4 := GenS1{F0: 65535, F1: 127}
    v5 := GenS1{F0: 744, F1: 1}
    _ = p0
    _ = v4
    _ = v5
    return 9223372036854775807
}

func main() {
    var v1 uint = 18446744073709551615
    var v2 uint16 = 32767
    var v3 uint = v1
    v6 := []int{}
    var v7 float32 = (-2.25 - 1.5)
    if false {
        v8 := GenS1{F0: v1, F1: 0}
        v1 += (^(v3))
        _ = v8
    } else {
        v9 := GenS1{F0: v3, F1: 1}
        _ = v9
    }
    v2 = (v2 / ((v2) | 1))
    if false {
        var v10 string = (("a" + "hello") + "")
        v2 = v2
        _ = v10
    } else {
        v2 = v2
        v3 = v3
        for fi11 := 0; fi11 < 13; fi11++ {
            if true {
                for fi12 := 0; fi12 < 11; fi12++ {
                    var v13 uint16 = (v2 / ((v2) | 1))
                    _ = fi12
                    _ = v13
                }
                v3 = 195
            }
            var v14 uint32 = uint32(v2)
            _ = fi11
            _ = v14
        }
    }
    fmt.Println(v2)
    if false {
        v6 = append(v6, -9223372036854775808)
        v3 -= uint(v2)
    }
    if len(v6) > 0 {
        v6[0] = 1
    }
    v3 = v3
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v1 -= (v3 | v3)
    v15 := GenS1{F0: v1, F1: -31}
    var v16 int16 = 182
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v6
    _ = v7
}
