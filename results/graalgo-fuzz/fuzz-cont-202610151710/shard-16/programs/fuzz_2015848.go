package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 uint }
func genFn0(p0 uint16, p1 int16) int64 {
    p0 -= (p0 >> ((42) & 63))
    fmt.Println(p0)
    p0 = (p0 ^ p0)
    for fi6 := 0; fi6 < 6; fi6++ {
        var v7 uint32 = 2
        v8 := []int{fi6}
        v9 := GenS1{F0: 600, F1: 826, F2: 255}
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    return -795
}

func main() {
    v1 := make(map[string]string)
    var v2 uint = 565
    var v3 int8 = 100
    var v4 float32 = 0.1
    var v5 int16 = 1000
    v4 -= float32(0.1)
    switch "foo bar" {
    case "a":
        if ((!true) || (-2.25 != v4)) {
            var v10 uint64 = 168
            if (!(!((v2 << ((v2) & 63)) != v2))) {
                delete(v1, "")
                if false {
                    var v11 int16 = (v5 & 32767)
                    delete(v1, "foo bar")
                    v3 = v3
                    _ = v11
                }
            }
            _ = v10
        }
    }
    if false {
        v5 += v5
        var v12 uint64 = 0
        _ = v12
    }
    v13 := []int{}
    v4 = (3.14159 / (float32((0.1 - 1.5)) * v4))
    for fi14 := 0; fi14 < 9; fi14++ {
        var v15 float32 = (v4 + v4)
        _ = fi14
        _ = v15
    }
    if len(v13) > 0 {
        v13[0] = 9223372036854775807
    }
    switch "foo bar" {
    case "a":
        _ = len(v1)
    }
    v4 = v4
    v16 := GenS1{F0: 0, F1: v2, F2: v2}
    var v17 int8 = v3
    v3 += -16
    fmt.Println(v5, v17)
    fmt.Println(v5)
    v2 -= v2
    for fi18 := 0; fi18 < 18; fi18++ {
        if len(v13) > 0 {
            v13[0] = 32767
        }
        v4 -= float32(0.0)
        _ = fi18
    }
    v19 := GenS1{F0: 42, F1: v2, F2: v2}
    var v20 uint64 = 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
