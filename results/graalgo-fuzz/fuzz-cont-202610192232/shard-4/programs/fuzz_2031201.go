package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8; F1 uint8 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 uint32) int16 {
    p0 = 0
    var v5 string = "a"
    switch ("" + "hello") {
    case "\n\t":
        v5 = v5
    case "a":
        v6 := GenS1{F0: v5}
        _ = v6
    }
    _ = p0
    _ = v5
    return -32768
}
func genFn1(p0 uint64, p1 uint) int8 {
    v7 := GenS2{F0: 100, F1: 52}
    var v8 uint64 = p0
    fmt.Println(v8, p0)
    v9 := []int{32767, -9223372036854775808, -411}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return -82
}

func main() {
    var v1 float32 = -2.25
    var v2 uint32 = 0
    var v3 int16 = -1
    v4 := make(map[int]uint32)
    var v10 int = 572
    if false {
        v11 := GenS3{F0: 127}
        _ = v11
    }
    v12 := GenS2{F0: -80, F1: 198}
    for fi13 := 0; fi13 < 10; fi13++ {
        var v14 float64 = -2.25
        var v15 int64 = 212
        for fi16 := 0; fi16 < 5; fi16++ {
            v17 := []int{}
            if (!false) {
                v3 -= (v3 / (((v3 ^ (v3 ^ (-(v3))))) | 1))
                var v18 float64 = float64((v1 - float32((1.5 / 1.5))))
                _ = v18
            }
            _ = fi16
            _ = v17
        }
        _ = fi13
        _ = v14
        _ = v15
    }
    for fi19 := 0; fi19 < 11; fi19++ {
        switch fi19 {
        case 9223372036854775807:
            v20 := GenS2{F0: -98, F1: 127}
            _ = v20
        default:
            v21 := []int{fi19, 9223372036854775807}
            _ = v21
        }
        for fi22 := 0; fi22 < 15; fi22++ {
            var v23 uint8 = 128
            v23 += (v23 << ((2) & 63))
            _ = fi22
            _ = v23
        }
        v10 += v10
        _ = fi19
    }
    var v24 bool = false
    v25 := GenS1{F0: ("" + "\n\t")}
    for fi26 := 0; fi26 < 2; fi26++ {
        if (0 > 32767) {
            var v27 uint32 = v2
            _ = v27
        }
        var v28 int32 = -1
        v3 = 0
        _ = fi26
        _ = v28
    }
    for fi29 := 0; fi29 < 18; fi29++ {
        v24 = true
        _ = fi29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
