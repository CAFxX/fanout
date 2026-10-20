package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 int8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 string }
func genFn0(p0 uint16, p1 int) int64 {
    v5 := GenS1{F0: 675, F1: 20, F2: 16}
    p0 += 278
    if true {
        fmt.Println(p0, p0)
        v6 := []int{p1}
        v7 := GenS1{F0: 0, F1: 167, F2: 1}
        _ = v6
        _ = v7
    } else {
        p1 = -387
        for fi8 := 0; fi8 < 1; fi8++ {
            var v9 int = fi8
            if true {
                for fi10 := 0; fi10 < 12; fi10++ {
                    v11 := GenS1{F0: 219, F1: 602, F2: 42}
                    _ = fi10
                    _ = v11
                }
            }
            fmt.Println(fi8)
            _ = fi8
            _ = v9
        }
        v12 := GenS1{F0: 4294967295, F1: 2, F2: 100}
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v5
    return -39
}

func main() {
    var v1 uint16 = 0
    var v2 uint8 = 112
    var v3 bool = true
    v4 := make(map[int]string)
    v3 = (!(true && (1 != 134)))
    if (0 >= (^(v1))) {
        v2 = v2
        for fi13 := 0; fi13 < 13; fi13++ {
            var v14 float64 = (float64((1e10 / 1e10)) * (((0.0 * float64(1e10)) + 1e10) / (1.5 + 0.1)))
            _ = fi13
            _ = v14
        }
        if ((0.0 + (1.5 / 1.0)) <= float64(float32((((0.1 / (0.0 + 1e10)) + 0.1) * (0.1 / 0.1))))) {
            fmt.Println(v2, v1)
        }
    } else {
        var v15 float64 = (1.5 / -2.25)
        v16 := GenS2{F0: "foo bar"}
        _ = v15
        _ = v16
    }
    v3 = (("hello" >= "a") || (true || v3))
    v17 := []int{}
    var v18 int16 = 857
    var v19 int32 = 2147483647
    v20 := GenS3{F0: ("a" + "hello")}
    for fi21 := 0; fi21 < 10; fi21++ {
        var v22 int16 = v18
        _ = fi21
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
