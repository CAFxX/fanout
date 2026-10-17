package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 uint; F1 string; F2 uint }
type GenS3 struct { F0 int8; F1 int32 }
type GenS4 struct { F0 uint32 }
func genFn0(p0 int64, p1 int16) uint16 {
    var v4 uint8 = 156
    v4 = v4
    var v5 int32 = -783
    v6 := []int{9223372036854775807}
    var v7 int8 = 127
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 186
}
func genFn1(p0 uint) int8 {
    for fi8 := 0; fi8 < 4; fi8++ {
        for fi9 := 0; fi9 < 8; fi9++ {
            p0 = p0
            _ = fi9
        }
        for fi10 := 0; fi10 < 5; fi10++ {
            v11 := GenS1{F0: 178, F1: 479}
            _ = fi10
            _ = v11
        }
        _ = fi8
    }
    switch -2147483648 {
    case -2147483648:
        var v12 int8 = 0
        _ = v12
    }
    fmt.Println(p0)
    for fi13 := 0; fi13 < 10; fi13++ {
        if false {
            fmt.Println(fi13)
            var v14 bool = (true && (!true))
            v14 = true
            _ = v14
        } else {
            var v15 float64 = 0.1
            _ = v15
        }
        _ = fi13
    }
    for fi16 := 0; fi16 < 11; fi16++ {
        var v17 uint = (^(p0))
        p0 += 898
        _ = fi16
        _ = v17
    }
    _ = p0
    return 0
}

func main() {
    var v1 float64 = -2.25
    v2 := make(map[int]bool)
    var v3 bool = false
    v1 = float64(float32((v1 / 1e10)))
    fmt.Println(v3)
    fmt.Println(v3, v3)
    v1 += (0.1 * 3.14159)
    for fi18 := 0; fi18 < 8; fi18++ {
        var v19 uint16 = 2
        _ = fi18
        _ = v19
    }
    v20 := GenS1{F0: 2, F1: 2147483647}
    if (!(!((42 <= 310) && (v1 > ((v1 / (float64(1.5) + float64((0.0 - 0.0)))) - (v1 + (v1 / (v1 - 0.0)))))))) {
        var v21 uint64 = 4294967295
        _ = v21
    } else {
        var v22 float64 = (float64(0.1) - (v1 / -2.25))
        _ = v22
    }
    v23 := GenS1{F0: 79, F1: 0}
    v24 := GenS2{F0: 4294967295, F1: "hello", F2: 788}
    var v25 int = 0
    v26 := GenS3{F0: -1, F1: 0}
    var v27 string = "foo bar"
    v28 := GenS4{F0: 290}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
}
