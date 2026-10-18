package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 bool }
func genFn0(p0 int, p1 int32, p2 int8) int8 {
    if false {
        v6 := []int{p0, (p0 + -1), -1}
        if false {
            p0 -= (p0 / (((^(p0))) | 1))
            for fi7 := 0; fi7 < 10; fi7++ {
                var v8 float64 = 0.0
                for fi9 := 0; fi9 < 19; fi9++ {
                    var v10 uint8 = 128
                    v6 = append(v6, fi9)
                    _ = fi9
                    _ = v10
                }
                _ = fi7
                _ = v8
            }
        }
        _ = v6
    }
    fmt.Println(p0, p1)
    v11 := GenS1{F0: -1, F1: "a"}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return p2
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[int]bool)
    v3 := make(map[string]bool)
    var v4 float32 = 1e10
    var v5 int16 = -1
    v12 := GenS2{F0: 65535}
    var v13 string = ("" + ("" + "hello"))
    for fi14 := 0; fi14 < 5; fi14++ {
        v5 = (^(v5))
        _ = fi14
    }
    fmt.Println(v4, v4)
    v15 := GenS3{F0: false}
    var v16 float64 = 0.0
    v17 := GenS3{F0: false}
    var v18 float64 = 1e10
    v18 = v16
    v13 = v13
    v4 = 0.1
    v16 = ((1e10 - 1.5) * float64((float32((float64((1e10 / (0.1 - float32(1e10)))) + v16)) * v4)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
