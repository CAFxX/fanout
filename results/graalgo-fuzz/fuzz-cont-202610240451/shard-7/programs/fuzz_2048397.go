package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 uint32 }
func genFn0(p0 uint8, p1 uint32) int64 {
    switch ("hello" + "a") {
    case "hello":
        if (65535 <= 0) {
            v6 := []int{9223372036854775807, 2147483647}
            _ = v6
        }
    case "a":
        var v7 bool = (227 >= 0)
        _ = v7
    default:
        if true {
            if false {
                for fi8 := 0; fi8 < 15; fi8++ {
                    for fi9 := 0; fi9 < 16; fi9++ {
                        p0 = (p0 * p0)
                        v10 := []int{fi8}
                        _ = fi9
                        _ = v10
                    }
                    _ = fi8
                }
                if (42 == 65535) {
                    fmt.Println(p0, p0)
                    fmt.Println(p0)
                }
                for fi11 := 0; fi11 < 13; fi11++ {
                    fmt.Println(p1)
                    var v12 int16 = 1
                    _ = fi11
                    _ = v12
                }
            }
            p1 = p1
            p1 = p1
        }
    }
    v13 := []int{-327, 0}
    _ = p0
    _ = p1
    _ = v13
    return -9223372036854775808
}

func main() {
    v1 := []string{}
    v2 := []string{"hello"}
    v3 := make(map[int]string)
    var v4 uint64 = 0
    v5 := make(map[string]int)
    v14 := GenS1{F0: 4294967295, F1: 255, F2: 65535}
    v4 += (v4 / ((0) | 1))
    var v15 float64 = float64(0.0)
    v15 += (v15 + 0.1)
    v4 = 227
    v16 := GenS1{F0: v4, F1: 42, F2: 65535}
    var v17 uint = 850
    var v18 uint16 = 42
    for fi19 := 0; fi19 < 7; fi19++ {
        if len(v2) > 0 {
            v2[0] = "hello"
        }
        _ = fi19
    }
    v4 = v4
    v18 = v18
    var v20 int16 = 32767
    var v21 float32 = ((3.14159 - 1e10) * 0.0)
    v17 -= v17
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
